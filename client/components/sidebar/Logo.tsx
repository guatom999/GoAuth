'use client'

import Image from "next/image"
import { useRouter } from "next/navigation"

const Logo = () => {
    const router = useRouter()

  return (
    <div
      className="flex flex-row gap-3 md:gap-0"
    >
        <Image alt="Logo" 
        className="
            flex 
            flex-row 
            items-center 
            justify-between 
            gap-3 
            rounded-full
            "
        height="100"
        width="40"
        src="/images/photo.png"
        >
        </Image>
    </div>
  )
}

export default Logo;