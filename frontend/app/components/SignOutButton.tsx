"use client";
import { useEffect } from "react";
import { signOut } from "../lib/auth";
import { firebaseAuth } from "../lib/firebase";





function SignInButton() {
          
    useEffect(() => {
        const unsubscribe = firebaseAuth.onAuthStateChanged((user) => {
            if (user) {
                console.log(user.email);
            } else {
                console.log("no user");
            }
        });
        return () => unsubscribe();
    }, []);

    return (
        <button onClick={signOut}>Logout</button>
    );
}
export default SignInButton;