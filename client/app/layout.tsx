import './globals.css'
import type { Metadata } from 'next'
import { Nunito } from "next/font/google"
import { Inter } from 'next/font/google'
import LoginModal from '../components/modal/LoginModal'
import RegisterModal from '../components/modal/RegisterModal'
import Navbar from '../components/navbar/Navbar'
import Button from '../components/Button'
import ClientOnly from '../components/ClientOnly'
import Modal from '../components/modal/Modal'

import getCurrentUser from '../components/actions/getCurrentUser'

import ToasterProvider from "./providers/ToastProviders"
import Providers from '../components/Providers'


const inter = Inter({ subsets: ['latin'] })

const font = Nunito({
  subsets: ["latin"]
})

export const metadata: Metadata = {
  title: 'SomeShitApp',
  description: 'SomeShitWebApp',
}

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {

  const currentUser = await getCurrentUser()

  // console.log(currentUser)

  return (
    <html lang="en">
      <body
      // className={inter.className}
      >
        {/* <Providers> */}
          <ToasterProvider />
          <LoginModal />
          <RegisterModal />
          <Navbar
          currentUser={currentUser}
          />
          {children}
        {/* </Providers> */}
      </body>
    </html>
  )
}
