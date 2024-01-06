import styles from "@/app/page.module.css";
import getRequestClient from "./lib/getRequestClient";
import { JobPreview } from "./components/Job";
import { cookies } from "next/headers";

export default async function Home() {

  const client = getRequestClient();
  const { jobs } = await client.jobs.Jobs()
  const isLoggedIn = cookies().has("auth-token");
  return (
    <div className="container mx-auto">
    <section className="">
      <h1 className="text-center text-6xl">Get Involved!</h1>

      <div>
       

        <div>
          {
            jobs.map((job, index) => (
              <JobPreview job={job} key={index} loggedIn={isLoggedIn}/>
            ))
          }
        </div>


     
      </div>
    </section>
    </div>
  
  );
}
