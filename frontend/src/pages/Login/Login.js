import React, { useState } from "react";
import "./Login.scss";

const Login = () => {
  const [fakeAuth, setAuth] = useState(false)
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const fakeLoginHandler = (event) => {
    if (email.length > 0 && password.length > 0){
      event.preventDefault()
      setAuth(true)
    }
  }
  return (
    <div className="hero-body">
      <div className="container has-text-centered">
        <div className="column is-4 is-offset-4">
          <div className="box">
            <p>{fakeAuth ? "Not Authorized": ""}</p>
            <form>
              <div className="field">
                <div className="control">
                  <input
                    className="input is-large"
                    type="email"
                    placeholder="Your Email"
                    autoFocus=""
                    onInput={(event) => { setEmail(event.target.value)}}
                  />
                </div>
              </div>

              <div className="field">
                <div className="control">
                  <input
                    className="input is-large"
                    type="password"
                    placeholder="Your Password"
                    onInput={(event) => { setPassword(event.target.value) }}
                  />
                </div>
              </div>
              <button disabled={email.length <= 0 || password.length <= 0} onClick={fakeLoginHandler} className="button is-block is-info is-large is-fullwidth">
                Login <i className="fa fa-sign-in" aria-hidden="true"></i>
              </button>
            </form>
          </div>
          <p>
            {/* <a className="has-text-white" href="#sign">
              Sign Up
            </a>
            &nbsp;·&nbsp; */}
            <a className="has-text-white" href="#pass">
              Forgot Password
            </a>
            {/* &nbsp;·&nbsp;
            <a className="has-text-white" href="#help">
              Need Help?
            </a> */}
          </p>
        </div>
      </div>
    </div>
  );
};

export default Login;
