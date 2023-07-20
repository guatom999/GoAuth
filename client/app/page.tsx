import Image from 'next/image'
import Navbar from '../components/navbar/Navbar'
import SideBar from '../components/sidebar/SideBar'
import Container from '../components/Container'
import DashBoard from '../components/dashboard/DashBoard'

export default function Home() {
  return (
    <div
      // className="
      // pt-24
      // grid 
      // grid-cols-1 
      // sm:grid-cols-2 
      // md:grid-cols-3 
      // lg:grid-cols-4
      // xl:grid-cols-5
      // 2xl:grid-cols-6
      // gap-8
      // "
      className="pt-24"
    >
      {/* <DashBoardContainer> */}
        <DashBoard/>
      {/* </DashBoardContainer> */}

    </div>


    // <main className="flex min-h-screen flex-col items-center justify-between p-24">
    //   {/* Hello Ja */}
    //   {/* <Navbar /> */}
    // </main>
  )
}
