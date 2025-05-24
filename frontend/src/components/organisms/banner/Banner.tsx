import React from 'react';
import { FaCheck } from 'react-icons/fa';

const Banner = () => {
  return (
    <div className="bg-gray-300 px-6 py-16 shadow-sm">
      <h1 className="text-3xl sm:text-3xl lg:text-3xl font-extrabold text-black ml-10">
        HANYA 10K PERBULAN BISA LIAT REKOMENDASI
      </h1>
      <h1 className="text-3xl sm:text-3xl lg:text-3xl font-extrabold text-black ml-10">
        BERDASARKAN CV KAMU
      </h1>

      <div className="flex flex-wrap items-center gap-8 justify-between ml-10">
        <div className="flex flex-wrap gap-8">
          <h3 className="text-xl sm:text-2xl lg:text-2xl font-extrabold text-black">
            <span className="flex items-center">
              50x pencocokan <FaCheck className="text-green-600 ml-2" />
            </span>
          </h3>
          <h3 className="text-xl sm:text-2xl lg:text-2xl font-extrabold text-black">
            <span className="flex items-center">
              Data tiap hari update <FaCheck className="text-green-600 ml-2" />
            </span>
          </h3>
        </div>

        <button
          className="text-2xl bg-white hover:bg-blue-900 text-black hover:text-white font-bold py-4 px-6 sm:px-8 lg:px-10 rounded-xl 
        transform transition-transform hover:translate-x-3 -translate-y-9 mr-20 focus:outline-none focus:ring-4 focus:ring-blue-300"
        >
          Ill Try It!
        </button>
      </div>
    </div>
  );
};

export default Banner;
