'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { createUserWithEmailAndPassword } from "firebase/auth";

import { firebaseAuth } from '../lib/firebase';

import {Container, Row, Col, Button, Form, FormGroup, Label, Input, Alert} from 'reactstrap';
import { signInWithGoogle } from '../lib/auth';

export default function SignUp() {
  const [email, setEmail] = useState("");
  const [passwordOne, setPasswordOne] = useState("");
  const [passwordTwo, setPasswordTwo] = useState("");
  const router = useRouter();
  const [error, setError] = useState(null);

  const onSubmit = event => {
    // setError(null)
    // //check if passwords match. If they do, create user in Firebase
    // // and redirect to your logged in page.
    // if(passwordOne === passwordTwo)
    //   createUserWithEmailAndPassword(firebaseAuth, email, passwordOne)
    //   .then(authUser => {
    //     console.log("Success. The user is created in Firebase")
    //     router.push("/logged_in");
    //   })
    //   .catch(error => {
    //     // An error occurred. Set error message to be displayed to user
    //     setError(error.message)
    //   });
    // else
    //   setError("Password do not match")
    // event.preventDefault();

    signInWithGoogle().then(authUser => {
            console.log("Success. The user is created in Firebase")
            console.log(authUser?.user.email)
            router.push("/logged_in");
          })
          .catch(error => {
            // An error occurred. Set error message to be displayed to user
            setError(error.message)
          });
    event.preventDefault();
  };

  return (
    <Container className="text-center custom-container">
      <Row>
        <Col>
          <Form 
            className="custom-form"
            onSubmit={onSubmit}>
          { error && <Alert color="danger">{error}</Alert>}
            <FormGroup row>
              <Label for="signUpEmail" sm={4}>Email</Label>
              <Col sm={8}>
                <Input
                  type="email"
                  value={email}
                  onChange={(event) => setEmail(event.target.value)}
                  name="email"
                  id="signUpEmail"
                  placeholder="Email" />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="signUpPassword" sm={4}>Password</Label>
              <Col sm={8}>
                <Input
                  type="password"
                  name="passwordOne"
                  value={passwordOne}
                  onChange={(event) => setPasswordOne(event.target.value)}
                  id="signUpPassword"
                  placeholder="Password" />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="signUpPassword2" sm={4}>Confirm Password</Label>
              <Col sm={8}>
                <Input
                  type="password"
                  name="password"
                  value={passwordTwo}
                  onChange={(event) => setPasswordTwo(event.target.value)}
                  id="signUpPassword2"
                  placeholder="Password" />
              </Col>
            </FormGroup>
            <FormGroup row>
             <Col>
               <Button>Sign Up</Button>
             </Col>
           </FormGroup>
          </Form>
        </Col>
      </Row>
    </Container>
  )
}
