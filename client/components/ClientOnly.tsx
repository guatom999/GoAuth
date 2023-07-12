'use client'

import React , { useState , useEffect } from 'react'

interface ClientOnlyProps {
    children: any
}

const ClientOnly: React.FC<ClientOnlyProps> = ({
    children
}) => {
    const [hasMounted, sethasMounted] = useState(false)

    useEffect(() => {
        sethasMounted(true)
    },[])

    if(!hasMounted){
        return null
    }
    return (
        <>
            {children}
        </>
    )
}

export default ClientOnly