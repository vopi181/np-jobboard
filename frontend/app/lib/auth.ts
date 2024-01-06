
import { initializeApp } from "firebase/app";
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";

import { firebaseAuth } from "./firebase";

// export function onAuthStateChanged(cb) {
//     return _onAuthStateChanged(firebaseAuth, cb);
// }

export async function signInWithGoogle() {
    const provider = new GoogleAuthProvider();

    try {
            return signInWithPopup(firebaseAuth, provider);
    } catch (error) {
            console.error("Error signing in with Google", error);
    }
}

export async function signOut() {
    try {
            return firebaseAuth.signOut();
    } catch (error) {
            console.error("Error signing out with Google", error);
    }
}