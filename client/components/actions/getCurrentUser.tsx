import { getServerSession } from "next-auth/next";

import { useSession } from 'next-auth/react'


import { authOptions } from "@/pages/api/auth/[...nextauth]";

export async function getSession(){
    return await getServerSession(authOptions)
}

export default async function getCurrentUser() {
    try{
        const session = await getSession()

        // const { data : session } = useSession()

        console.log("session is" , session)

        // if(!session?.user?.email){
        //     return null
        // }

        // // const currentUser = await prisma.user.findUnique({
        // //     where:{
        // //         email:session.user.email as string
        // //     }
        // // })

        // if(!currentUser){
        //     return null
        // }

        return session

        // return {
        //     ...currentUser,
        //     createdAt: currentUser.createdAt.toISOString(),
        //     updatedAt: currentUser.updatedAt.toISOString(),
        //     emailVerified: currentUser.emailVerified?.toISOString() || null,
        // };
            
    }catch(error:any){
        return null
    }
}