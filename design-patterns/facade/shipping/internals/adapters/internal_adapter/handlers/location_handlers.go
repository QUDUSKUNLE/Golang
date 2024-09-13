package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var wg sync.WaitGroup

// @Summary Submit addresses
// @Description create addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param Body body domain.LocationDto true "Create addresses"
// @Param Authorization header string true "Bearer token"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [post]
func (handler *HTTPHandler) PostAddress(context echo.Context) error {
	location := new(domain.LocationDto)
	if err := ValidateStruct(context, location); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err,
			context)
	}

	// Parse user
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}

	// Make call to internal adapter to save register
	location.UserID = user.ID
	err = handler.internalServicesAdapter.NewLocationAdaptor(*location);
	if err != nil {
		return ComputeErrorResponse(http.StatusConflict, ADDRESS_ALREADY_EXIST, context)
	}

	// Need to run this with goroutine, working on this
	var addressSync sync.WaitGroup
	var externalAddress map[string]interface{}
	var terminalLocation *domain.Location
	terminalLocationChannel := make(chan domain.Location)

	// InternalTerminalAdaptor function
	internalTerminalAdaptor := func (addSync *sync.WaitGroup, location domain.Location) {
		defer addSync.Done()
		err := handler.internalServicesAdapter.TerminalUpdateAddressAdaptor(location)
		if err != nil {
			panic(err)
		}
	}
	// ExternalTerminalAdaptor function
	externalTerminalAdaptor := func (addSync *sync.WaitGroup, address domain.Address, index int, location domain.LocationDto, locationChannel chan <- domain.Location)   {
		defer addSync.Done()
		externalAddress, _ = handler.externalServicesAdapter.TerminalCreateAddressAdaptor(address)
		if externalAddress["data"] != nil {
			result := externalAddress["data"].(map[string]interface{})
			address_id := result["address_id"].(string)
			location.Address[index].TerminalAddressID = address_id
			terminalLocation = &domain.Location{
				TerminalAddressID: address_id,
				Address: location.Address[index],
				Description: location.Address[index].Description,
				UserID: location.UserID,
			}
			// Send terminalLocation to the channel
			locationChannel <- *terminalLocation
		}
	}
	// Close channel function
	closeTerminalChannel := func(terminalGroup *sync.WaitGroup, channel chan <- domain.Location) {
		terminalGroup.Wait()
		close(channel)
	}
	// Goroutine externalTerminalAdaptor
	for index, address := range location.Address {
		addressSync.Add(1)
		go externalTerminalAdaptor(&addressSync, address, index, *location, terminalLocationChannel)
	}
	// Goroutine close channel
	go closeTerminalChannel(&addressSync, terminalLocationChannel)

	// Goroutine internalTerminalAdaptor
	for terminalChannel := range terminalLocationChannel {
		addressSync.Add(1)
		go internalTerminalAdaptor(&addressSync, terminalChannel)
	}
	addressSync.Wait()

	// Process valid location data
	return ComputeResponseMessage(http.StatusCreated, ADDRESSES_SUBMITTED_SUCCESSFULLY, context)
}

// @Summary Get a address
// @Description get a address
// @Tags Address
// @Accept json
// @Produce json
// @Param addressID path string true "Address ID"
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses/{addressID} [get]
func (handler *HTTPHandler) GetAddress(context echo.Context) error {
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return err
	}

	ID := context.Param("addressID")
	addressID, err := uuid.Parse(ID)
	if err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	location, err := handler.internalServicesAdapter.GetLocationAdaptor(addressID, user.ID);
	if err != nil {
		return ComputeErrorResponse(http.StatusConflict, err.Error(), context)
	}
	return ComputeResponseMessage(http.StatusOK, location, context)
}

// @Summary Get addresses
// @Description get addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param description query string false "Description"
// @Failure 400 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /addresses [get]
func (handler *HTTPHandler) GetAddresses(context echo.Context) error {
	user, err := PrivateMiddlewareContext(context, string(domain.USER))
	if err != nil {
		return ComputeErrorResponse(http.StatusUnauthorized, UNAUTHORIZED_TO_PERFORM_OPERATION, context)
	}
	description := context.QueryParam("description")
	if description == "" {
		locations, err := handler.internalServicesAdapter.GetLocationsAdaptor(user.ID);
		if err != nil {
			return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
		}
		return ComputeResponseMessage(http.StatusOK, locations, context)
	} else {
		location, err := handler.internalServicesAdapter.QueryLocationAdaptor(user.ID, description);
			if err != nil {
				return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
			}
			return ComputeResponseMessage(http.StatusOK, location, context)
	}
}

// Example of LiveLock
func (handler *HTTPHandler) LiveLock(context echo.Context) error {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Microsecond) {
			cadence.Broadcast()
		}
	}()
	takeStep := func(){
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}
	tryDir := func (dirName string, dir *int32, out *bytes.Buffer) bool  {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}
	var left, right int32
	tryLeft := func (out *bytes.Buffer) bool { return tryDir("left", &left, out)}
	tryRight := func (out *bytes.Buffer) bool { return tryDir("right", &right, out)}

	walk := func (walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String())}()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation", name)
	}
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
	return ComputeResponseMessage(http.StatusOK, "LiveLock", context)
}

// Example of DeadLock
func (handler *HTTPHandler) DeadLock(context echo.Context) error {
	type Value struct {
		mu sync.Mutex // guards
		value int
	}
	printSum := func(v1, v2 *Value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		// time.Sleep(2 *time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum = %v\n", v1.value + v2.value)
	}

	var a, b Value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()

	return ComputeResponseMessage(http.StatusOK, "DeadLock", context)
}

// Example of DeadLock
func (handler *HTTPHandler) Starvation(context echo.Context) error {
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops.\n", count)
	}
	politeWorker := func ()  {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}
	niceWorker := func ()  {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("Nice worker was able to execute %v work loops.\n", count)
	}
	wg.Add(3)
	go greedyWorker()
	go politeWorker()
	go niceWorker()
	wg.Wait()
	return ComputeResponseMessage(http.StatusOK, "Starvation", context)
}

func (handler *HTTPHandler) Example(context echo.Context) error {
	salutation := "Hello"
	sayHello := func() {
		defer wg.Done()
		salutation = fmt.Sprintf("%s, you are welcome.", salutation)
	}
	salutes := func(salute string) {
		defer wg.Done()
		fmt.Println(salute)
	}
	for _, salute := range []string{"Hello", "greetings", "good day"} {
		wg.Add(1)
		go salutes(salute)
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
	fmt.Println(salutation)
	return ComputeResponseMessage(http.StatusOK, "Example", context)
}
