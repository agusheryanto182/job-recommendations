import React from 'react';
import Navbar from '../components/Navbar';

const Homepage = () => {
  return (
    <div className="min-h-screen bg-white">
      {/* Navbar Section */}
      <Navbar />

      {/* Hero Section */}
      <div className="container mx-auto px-6 py-20 flex flex-col lg:flex-row items-center justify-between gap-12">
        {/* Left Text Section */}
        <div className="flex flex-col gap-6 max-w-xl">
          <h1 className="text-4xl sm:text-6xl font-extrabold text-black">
            JobMatcher <br />
            <span className="text-black">for Your Life.</span>
          </h1>
          <p className="text-gray-700 leading-relaxed">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </p>

          {/* Statistik Section */}
          <div className="flex flex-wrap gap-4 pt-4">
            <div className="border rounded-xl bg-gray-300 px-6 py-4 text-center text-black">
              <p className="text-lg font-bold">10.000</p>
              <p className="text-sm">PENGGUNA</p>
            </div>
            <div className="border rounded-xl bg-gray-300 px-6 py-4 text-center text-black">
              <p className="text-lg font-bold">10.000</p>
              <p className="text-sm">JUMLAH KLIK</p>
            </div>
            <div className="border rounded-xl bg-gray-300 px-6 py-4 text-center text-black">
              <p className="text-lg font-bold">10.000</p>
              <p className="text-sm">CV MATCHED</p>
            </div>
          </div>
        </div>

        {/* Login Section */}
        <div className="bg-gray-300 rounded-3xl px-10 py-20 w-full max-w-lg h-[500px] text-center flex items-center justify-center">
          <button
            className="bg-white text-black font-semibold max-w-xs px-6 py-4 rounded-xl border hover:bg-gray-100 w-full"
          >
            Masuk dengan Google
          </button>
        </div>
      </div>
    </div>
  );
};

export default Homepage;