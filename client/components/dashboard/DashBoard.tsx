import React from 'react'
import Image from "next/image";
import { getProductsPrice } from '../../services/product'

// interface DashBoardProps {
const src = "https://f.btwcdn.com/store-49336/product/489e6105-7130-7dc0-b3ff-641823ca656b.jpg"
const alt = "Bossza"
// }

const DashBoard: React.FC = async () => {

  const totalPrice = await getProductsPrice({limit : 10})

  console.log(totalPrice)

  return (
    <div className='grid xl:grid-cols-5 gap-4 p-4 pl-24'>
      <div className='lg:col-span-2 col-span-1 bg-white flex justify-between w-full border p-4 rounded-lg'>
        <div className='flex flex-col w-full pb-4'>
          <p className='text-2xl font-bold'>{totalPrice}</p>
          <p className='text-gray-600'>Total Products Price</p>
        </div>
        <p className='bg-green-200 flex justify-center items-center p-2 rounded-lg'>
          <span className='text-green-700 text-lg'>+18%</span>
        </p>
      </div>
      <div className='lg:col-span-2 col-span-1 bg-white flex justify-between w-full border p-4 rounded-lg'>
        <div className='flex flex-col w-full pb-4'>
          <p className='text-2xl font-bold'>$1,437,876</p>
          <p className='text-gray-600'>YTD Revenue</p>
        </div>
        <p className='bg-green-200 flex justify-center items-center p-2 rounded-lg'>
          <span className='text-green-700 text-lg'>+11%</span>
        </p>
      </div>
      <div className='bg-white flex justify-between w-full border p-4 rounded-lg'>
        <div className='flex flex-col w-full pb-4'>
          <p className='text-2xl font-bold'>11,437</p>
          <p className='text-gray-600'>Customers</p>
        </div>
        <p className='bg-green-200 flex justify-center items-center p-2 rounded-lg'>
          <span className='text-green-700 text-lg'>+17%</span>
        </p>
      </div>
    </div>
    // <div
    //   className='flex w-full md:col-span-2 relative lg:h-[70vh] h-[50vh] m-auto p-4 border rounded-lg '
    // >
    //   <Image
    //     className="w-full"
    //     src={src}
    //     alt={alt}
    //     width={0}
    //     height={0}
    //     sizes="100"
    //   />
    // </div>
  )
}

export default DashBoard