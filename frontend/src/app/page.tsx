import React from 'react'

const Homepage = () => {
  return (
    <div className='container mx-auto flex flex-col p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-poppins)]'>
      <div>
      <h1 className='text-3xl font-bold sm:text-5xl tracking-light text-orange-600'>JobMatcher</h1>
      <h2 className='text-2xl font-bold sm:text-4xl tracking-light text-white'>For Your Life</h2>
      <p className='pt-4 text-white leading-relaxed'>Lorem ipsum dolor sit amet, consectetur adipisicing elit. 
        Earum nihil odit aliquid, sit illum repellat excepturi veritatis iste obcaecati ipsa, 
        possimus rerum perspiciatis unde laboriosam non placeat numquam est voluptatibus.
      </p>
        <div className='flex gap-4 pt-8'>
          <button className='px-8 py-4 text-white bg-orange-600 rounded-lg bg-opacity-80 font-medium 
          hover:bg-white hover:text-orange-600'>Lorem</button>
          <button className='px-8 py-4 text-orange-600 bg-white rounded-lg bg-opacity-80 font-medium 
          hover:bg-orange-600 hover:text-white'>Ipsum</button>
          <button className='px-8 py-4 text-orange-600 bg-white rounded-lg bg-opacity-80 font-medium 
          hover:bg-orange-600 hover:text-white'>dolor</button>
        </div>
      </div>
    </div>
  )
}

export default Homepage