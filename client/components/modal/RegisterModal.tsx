'use client'

import axios from 'axios'
import { AiFillGithub } from 'react-icons/ai'
import { FcGoogle } from 'react-icons/fc'
import { useCallback, useState } from 'react'
import {
    FieldValues,
    SubmitHandler,
    useForm
} from "react-hook-form";

import { toast } from 'react-hot-toast'

import useRegisterModal from '@/app/hooks/useRegisterModal'
import Modal from './Modal'
import Heading from '../Heading'
import Input from '../inputs/Input'
import Button from '../Button'
import { signIn } from 'next-auth/react'

const RegisterModal : React.FC= () => {
    const registerModal = useRegisterModal();
    const [isLoading, setIsLoading] = useState(false);

    const { register, handleSubmit, formState: { errors } } = useForm<FieldValues>(
    //     {
    //     defaultValues: {
    //         name: '',
    //         email: '',
    //         password: '',
    //     }
    // }
    );

    const onSubmit: SubmitHandler<FieldValues> = (data) => {
        setIsLoading(true);

        axios.post('http://localhost:8000/signup', data).then((result) => {
            // toast.success("Create Account Success") // Not working
            registerModal.onClose();
        }).catch((error) => {
            alert(error)
            // toast.error("Something Went Wrong") // Not Working
        }).finally(() => {
            // toast.success("Create Account Success") // Not Working
            registerModal.onClose();
            console.log("Finally")
            setIsLoading(false)
        })

    }

    const bodyContent = (
        <div className="flex flex-col gap-4">
            <Heading
                title="Create Account"
                // subtitle="Create an account!"
            />
            {/* <Input
                id="email"
                label="Email"
                disabled={isLoading}
                register={register}
                errors={errors}
                required
            /> */}
            <Input
                id="username"
                label="Name"
                disabled={isLoading}
                register={register}
                errors={errors}
                required
            />
            <Input
                id="password"
                label="Password"
                type="password"
                disabled={isLoading}
                register={register}
                errors={errors}
                required
            />
        </div>
    )

    // const footerContent = (
    //     <div
    //         className='flex flex-col gap-4 mt-3'
    //     >
    //         <hr />
    //         <Button
    //             outline
    //             label="Continue with Google"
    //             icon={FcGoogle}
    //             onClick={() => { }}
    //         />
    //         <Button
    //             outline
    //             label="Continue with Github"
    //             icon={AiFillGithub}
    //             onClick={() => signIn('github')}
    //         />
    //         <div
    //             className='
    //                 text-neutral-500
    //                 text-center
    //                 mt-4
    //                 font-light
    //             '
    //         >
    //             <div className='justify-center flex flex-row  items-center text-center gap-2'>
    //                 <div>
    //                     Already have an account
    //                 </div>
    //                 <div
    //                     onClick={registerModal.onClose}
    //                     className='
    //                     text-neutral-800
    //                     cursor-pointer
    //                     hover:underline
    //                     font-bold    
    //                     '
    //                 >
    //                     Login
    //                 </div>

    //             </div>

    //         </div>
    //     </div>
    // )

    return (
        <Modal
            disabled={isLoading}
            isOpen={registerModal.isOpen}
            title="Register"
            actionLabel="Continue"
            onClose={registerModal.onClose}
            onSubmit={handleSubmit(onSubmit)}
            body={bodyContent}
            // footer={footerContent}
        />
    );
}

export default RegisterModal