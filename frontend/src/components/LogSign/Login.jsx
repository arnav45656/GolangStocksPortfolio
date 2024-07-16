import React from 'react';
import {
  MDBBtn,
  MDBContainer,
  MDBCard,
  MDBCardBody,
  MDBCardImage,
  MDBRow,
  MDBCol,
  MDBInput,
  MDBCheckbox
}
from 'mdb-react-ui-kit';

import { useState } from 'react';

import './Login.css'

function Login() {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = async () => {
        try {
            const response = await axios.post('http://localhost:8080/login', { username, password });
            const token = response.data.token;
            setAuthToken(token);
            localStorage.setItem('jwtToken', token);
            // Redirect to home page or another protected page
        } catch (error) {
            console.error('Login error', error);
        }
    };



  return (
    <MDBContainer className='my-5'>
      <MDBCard>

        <MDBRow className='g-0 d-flex align-items-center'>

          <MDBCol md='4'>
            <MDBCardImage src='https://mdbootstrap.com/img/new/ecommerce/vertical/004.jpg' alt='phone' className='rounded-t-5 rounded-tr-lg-0' fluid />
          </MDBCol>

          <MDBCol md='8'>

            <MDBCardBody>

              <MDBInput wrapperClass='mb-4' label='Username' id='form1' type='text'/>
              <MDBInput wrapperClass='mb-4' label='Password' id='form2' type='password'/>

              

              <MDBBtn onClick={handleLogin} className="mb-4 w-100">Sign in</MDBBtn>

            </MDBCardBody>

          </MDBCol>

        </MDBRow>

      </MDBCard>
    </MDBContainer>
  );
}

export default Login;