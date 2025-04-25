import React, { ReactNode } from "react";
import Head from "next/head";

interface LayoutProps {
  children: ReactNode;
  title?: string;
}

const Layout = ({ children, title = "Hospital Diagnostic WebApp" }: LayoutProps) => {
  return (
    <>
      <Head>
        <title>{title}</title>
        <meta name="description" content="Web app for hospitals, diagnostic centres, and patients" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <html lang="en">
        <body className="bg-secondary text-primary">
          {/* Header Section */}
          <header className="w-full p-4 bg-skyblue text-center text-white font-bold">
            {title}
          </header>
          
          {/* Main Content */}
          <main className="p-6 flex flex-col items-center justify-center min-h-screen">
            {children}
          </main>
          
          {/* Footer Section */}
          <footer className="w-full p-4 bg-skyblue text-center text-white">
            &copy; {new Date().getFullYear()} Hospital Diagnostic WebApp
          </footer>
        </body>
      </html>
    </>
  );
};

export default Layout;
