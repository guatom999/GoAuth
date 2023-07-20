import React from 'react';
import Link from 'next/link';
import Image from 'next/image';
import { RxSketchLogo, RxDashboard, RxPerson } from 'react-icons/rx';
import { HiOutlineShoppingBag } from 'react-icons/hi';
import { FiSettings } from 'react-icons/fi';
import Logo from './Logo'

interface SideBarProps {
  children?: any
}


const Sidebar: React.FC<SideBarProps> = ({ children }) => {
  return (
    <div className="fixed">
      <div className="w-20 h-screen p-4 border-r-[1px] flex flex-col">
        <div className='flex flex-col items-center'>
          <Link href='/'>
            <div className=' hover:bg-gray-200 hover:rounded-3xl cursor-pointer my-4 p-3 inline-block'>
              <RxDashboard size={20} />
            </div>
          </Link>
        </div>
        <div className='flex flex-col items-center'>
          <Link href='/'>
            <div className=' hover:bg-gray-200 hover:rounded-3xl cursor-pointer my-4 p-3 inline-block'>
              <RxDashboard size={20} />
            </div>
          </Link>
        </div>
        <div className='flex flex-col items-center'>
          <Link href='/'>
            <div className=' hover:bg-gray-200 hover:rounded-3xl cursor-pointer my-4 p-3 inline-block'>
              <RxDashboard size={20} />
            </div>
          </Link>
        </div>
        <div className='flex flex-col items-center'>
          <Link href='/'>
            <div className=' hover:bg-gray-200 hover:rounded-3xl cursor-pointer my-4 p-3 inline-block'>
              <RxDashboard size={20} />
            </div>
          </Link>
        </div>
      </div>
    </div>
    // <div className='fixed'>
    //   <div className='w-20 h-screen p-4 bg-white border-r-[1px]'>
    //     <div className='flex flex-col items-center'>
    //       <Link href='/'>
    //         <div className='bg-gray-100 hover:bg-gray-200 cursor-pointer my-4 p-3 rounded-lg inline-block'>
    //           <RxDashboard size={20} />
    //         </div>
    //       </Link>
    //       <Link href='/customers'>
    //         <div className='bg-gray-100 hover:bg-gray-200 cursor-pointer my-4 p-3 rounded-lg inline-block'>
    //           <RxPerson size={20} />
    //         </div>
    //       </Link>
    //       <Link href='/orders'>
    //         <div className='bg-gray-100 hover:bg-gray-200 cursor-pointer my-4 p-3 rounded-lg inline-block'>
    //           <HiOutlineShoppingBag size={20} />
    //         </div>
    //       </Link>
    //       <Link href='/'>
    //         <div className='bg-gray-100 hover:bg-gray-200 cursor-pointer my-4 p-3 rounded-lg inline-block'>
    //           <FiSettings size={20} />
    //         </div>
    //       </Link>
    //     </div>
    //   </div>
    //   <main className='ml-20 w-full'>{children}</main>
    // </div>
  );
};

export default Sidebar;