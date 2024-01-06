// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getAuth } from "firebase/auth";


// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyB9uQoIW115qykqefDAzDxuG3pooj3tBVE",
  authDomain: "np-jobboard.firebaseapp.com",
  projectId: "np-jobboard",
  storageBucket: "np-jobboard.appspot.com",
  messagingSenderId: "694616445718",
  appId: "1:694616445718:web:da6bb8e296b9ad2823f637",
  measurementId: "G-FTRNT73MN6"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
// const analytics = getAnalytics(app);

export const firebaseAuth = getAuth(app)
export default app;