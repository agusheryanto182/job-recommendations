import React from 'react';
import { FaUserCircle } from 'react-icons/fa';

const Navbar = () => {
  return (
    <nav className="bg-gray-300 px-6 py-4 shadow-sm">
      <div className="container mx-auto flex justify-between items-center">
        <div className="text-xl font-bold text-white">Job<span className="text-black">Matcher</span></div>
        <ul className="flex gap-40 justify-center font-semibold text-black">
          <li className="hover:text-white flex items-center">HOME</li>
          <li className="hover:text-white flex items-center">HASIL</li>
          <li className="hover:text-white flex items-center">PAKET</li>
        </ul>
        <div className="text-2xl text-gray-700">
          <FaUserCircle />
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
