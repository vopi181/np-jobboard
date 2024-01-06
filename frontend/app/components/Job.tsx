'use client';

import { useRouter } from "next/router";
import { FormEvent, useEffect } from "react";
import styles from "@/app/page.module.css";
import Image from "next/image"
import Link from "next/link";

interface JobPreviewProps {
    job: {
        id: string;
        organization: string;
        title: string;
        is_paid: boolean;
        city: string;
        img: string;
        posted: string;
        description: string;
    },
    loggedIn: boolean;
}


export function JobPreview(props: JobPreviewProps) {
    const { job } = props;
    const { id, organization, title, is_paid, city, img, posted, description } = job;
    const loggedIn  = props.loggedIn;
   
    
    var onViewDetails = async (event: any) => {
        alert("Placeholder Data")
    }
  

    return (

        <div className="container mx-auto transition ease-in-out hover:bg-stone-500 rounded my-5 p-1 bg-stone-300 hover:drop-shadow-2xl drop-shadow-xl">

            <div className="grid grid-col-2  m-2">
                <h2 className="text-xl">{title}</h2>
                <div>
                    <div className="flex justify-end ">
                        <p className="text-xl text-right">{organization}</p>

                        <Image className="drop-shadow-xl" src={img} alt="job image" width={100} height={100} />
                    </div>
                </div>

                <p className="text-lg opacity-80">{city}</p>
                <p className="text-lg">{
                    
                    new Date(posted).toLocaleDateString("en-US", {
                        month: "long",
                        day: "numeric",
                        year: "numeric",
                    })
                }</p>

                {is_paid ? (<p className="text-lg">Paid ✔️</p>) : (<p className="text-lg">Unpaid ❌</p>)}

                <h3 className="text-md">{description}</h3>

                <div className="flex justify-end space-x-2">
                    

                    
                        <button onClick={onViewDetails} className="bg-white text-black border-none py-1 px-2 rounded cursor-pointer text-lg drop-shadow">View Details</button>
                    
                    
                </div>
            </div>
        </div>

    );
}

