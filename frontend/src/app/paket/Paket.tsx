import React from 'react'

const Paket = () => {
  return (
    <div className="flex flex-col items-center min-h-screen mt-40 px-4">
      <h1 className="text-4xl font-bold mb-12 text-black">Pilihan Paket</h1>
      <div className="flex flex-wrap justify-center gap-8">
        {/* Free Plan */}
        <div className="bg-gray-300 rounded-3xl px-10 py-20 w-full md:w-[30rem] h-[600px] text-black">
          <h1 className="text-3xl font-bold mb-6">Free Plan</h1>
          <ul className="list-disc list-inside text-black text-lg">
            <li>1x Pencocokan</li>
            <li>Mendapatkan 15 rekomendasi pekerjaan</li>
          </ul>
          <div className="flex justify-center mt-70">
            <button
              type="button"
              className="text-xl bg-white hover:bg-blue-900 text-black hover:text-white font-bold py-4 px-10 rounded-xl">
              Coba Sekarang
            </button>
          </div>
        </div>

        {/* Premium Plan */}
        <div className="bg-gray-300 rounded-3xl px-10 py-20 w-full md:w-[30rem] h-[600px] text-black">
          <h1 className="text-3xl font-bold mb-6">Premium Plan</h1>
          <ul className="list-disc list-inside text-black text-lg">
            <li>50x Pencocokan</li>
            <li>Mendapatkan 10 rekomendasi pekerjaan</li>
            <li>Score kecocokan di setiap pekerjaan</li>
          </ul>
          <div className="flex justify-center mt-63">
            <button
              type="button"
              className="text-xl bg-white hover:bg-blue-900 text-black hover:text-white font-bold py-4 px-10 rounded-xl">
              Coba Sekarang
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Paket
