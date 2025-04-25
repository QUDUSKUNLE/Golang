import Link from "next/link";
import Layout from "./layout";

export default function HomePage() {
  return (
    <Layout title="Welcome to Hospital Diagnostic WebApp">
      <div className="flex flex-col items-center text-center space-y-6">
        {/* Header Section */}
        <h1 className="text-4xl font-bold text-primary">Hospital Diagnostic WebApp</h1>
        <p className="text-lg text-primary">
          Manage patients, doctors, and diagnostic centres efficiently!
        </p>

        {/* Login and Signup Options */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 w-full max-w-4xl">
          {/* Patients Section */}
          <div className="bg-white shadow-lg rounded-lg p-6">
            <h2 className="text-xl font-bold text-skyblue mb-4">Patients</h2>
            <p className="text-primary mb-4">
              Access your health records and book appointments.
            </p>
            <div className="flex justify-between">
              <Link href="/patients/login">
                <button className="bg-skyblue text-white font-bold py-2 px-4 rounded hover:bg-primary hover:text-white transition">
                  Login
                </button>
              </Link>
              <Link href="/patients/signup">
                <button className="border border-skyblue text-skyblue font-bold py-2 px-4 rounded hover:bg-skyblue hover:text-white transition">
                  Signup
                </button>
              </Link>
            </div>
          </div>

          {/* Doctors Section */}
          <div className="bg-white shadow-lg rounded-lg p-6">
            <h2 className="text-xl font-bold text-skyblue mb-4">Doctors</h2>
            <p className="text-primary mb-4">
              Manage your appointments and consult with patients.
            </p>
            <div className="flex justify-between">
              <Link href="/doctors/login">
                <button className="bg-skyblue text-white font-bold py-2 px-4 rounded hover:bg-primary hover:text-white transition">
                  Login
                </button>
              </Link>
              <Link href="/doctors/signup">
                <button className="border border-skyblue text-skyblue font-bold py-2 px-4 rounded hover:bg-skyblue hover:text-white transition">
                  Signup
                </button>
              </Link>
            </div>
          </div>

          {/* Diagnostic Centres Section */}
          <div className="bg-white shadow-lg rounded-lg p-6">
            <h2 className="text-xl font-bold text-skyblue mb-4">Diagnostic Centres</h2>
            <p className="text-primary mb-4">
              Manage test records and collaborate with patients.
            </p>
            <div className="flex justify-between">
              <Link href="/diagnostic/login">
                <button className="bg-skyblue text-white font-bold py-2 px-4 rounded hover:bg-primary hover:text-white transition">
                  Login
                </button>
              </Link>
              <Link href="/diagnostic/signup">
                <button className="border border-skyblue text-skyblue font-bold py-2 px-4 rounded hover:bg-skyblue hover:text-white transition">
                  Signup
                </button>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
}
