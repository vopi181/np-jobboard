"use client";
import { doc } from "firebase/firestore";
import { signInWithGoogle, signOut } from "../lib/auth";
import { on } from "events";
import { FormEvent } from "react";

interface SignInButtonProps {
    loggedin: boolean;
}

function SignInButton(props: SignInButtonProps) {

    var onSigninSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        // debugger;

        await signInWithGoogle().then(authUser => {
            authUser?.user.getIdToken().then(token => {
                let tokenInput = document.getElementById("token-inp") as HTMLInputElement;
                tokenInput.value = token;
                let form = document.getElementById("token-form") as HTMLFormElement;
                form.setAttribute("action", "/auth/login");
                form.setAttribute("method", "POST");

                form.submit();


            })
        })
    }

    var onSignOutSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        // debugger;

        let form = document.getElementById("token-form") as HTMLFormElement;
        form.setAttribute("action", "/auth/logout");
        form.setAttribute("method", "POST");

        form.submit();
    }

    return (

        <div>
            {props.loggedin ?
                <form onSubmit={onSignOutSubmit} id="token-form">
                    <input type="hidden" name="token" id="token-inp" />
                    <button type="submit">Logout</button>
                </form>

                :

                <form onSubmit={onSigninSubmit} id="token-form">
                    <input type="hidden" name="token" id="token-inp" />
                    <button type="submit">Login with Google</button>
                </form>
            }
        </div>





    );
}
export default SignInButton;