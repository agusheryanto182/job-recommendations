'use client'

import React from 'react'
import { FaUserCircle } from 'react-icons/fa'
import Link from 'next/link'

const Navbar = () => {
  return (
    <nav className="fixed top-0 left-0 w-full z-50 bg-gray-300 bg-opacity-90 backdrop-blur-md px-6 py-4 shadow-md">
      <div className="container mx-auto flex flex-col md:flex-row justify-between items-center gap-4">
        <Link href="/" className="text-xl font-bold text-white">
          Job<span className="text-black">Matcher</span>
        </Link>
        <ul className="flex flex-col md:flex-row gap-6 md:gap-40 text-black font-semibold items-center">
          <li>
            <Link href="/" className="hover:text-white">HOME</Link>
          </li>
          <li>
            <Link href="/hasil" className="hover:text-white">HASIL</Link>
          </li>
          <li>
            <Link href="/paket" className="hover:text-white">PAKET</Link>
          </li>
        </ul>
        <div className="text-2xl text-gray-700">
          <Link href="/user">
            <FaUserCircle />
          </Link>
        </div>
      </div>
    </nav>
  )
}

export default Navbar
