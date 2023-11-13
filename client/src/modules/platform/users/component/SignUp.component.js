import * as React from "react";

export default function SignUp() {
    return (
        <>
            <div id="login-page" className="container-lg py-3">
                <div className="px-1 py-1 bg-white row align-items-center">
                    <div className="row">
                        <div className="col-lg-2"></div>
                        <div className="col-lg-4 text-start my-auto">
                            <img src="../images/logo.png" alt="Logo"></img>
                            <p className="login-page-left-heading text-info m-0 px-3 py-2">
                                KANBAN
                            </p>
                        </div>
                        <div className="col-lg-4 my-auto login-page-right">
                            <div className="row  ">
                                <div className="col-lg-12 text-center mb-3 ">
                                    <img
                                        className="login-page-right-image"
                                        src="../images/logo.png"
                                        alt="Logo"
                                    ></img>
                                    <p className="login-page-right-heading  text-dark m-0 px-3 py-2">
                                        Create an account
                                    </p>
                                    <p className="login-page-right-paragraph  text-secondary m-0 px-3 py-2">
                                        Start your 30-day free trial.
                                    </p>
                                </div>

                                <div className="row">
                                    <div className="col-lg-12 p-0">
                                        <form>
                                            <div className="mb-3">
                                                <label
                                                    for="exampleInputEmail1"
                                                    className="form-label login-page-right-form-label"
                                                >
                                                    Name*
                                                </label>
                                                <input
                                                    type="email"
                                                    className="form-control login-page-right-form-input"
                                                    id="exampleInputEmail1"
                                                    aria-describedby="emailHelp"
                                                    placeholder="Enter your name"
                                                    required
                                                />
                                            </div>
                                            <div className="mb-3">
                                                <label
                                                    for="exampleInputEmail1"
                                                    className="form-label login-page-right-form-label"
                                                >
                                                    Email*
                                                </label>
                                                <input
                                                    type="email"
                                                    className="form-control login-page-right-form-input"
                                                    id="exampleInputEmail1"
                                                    aria-describedby="emailHelp"
                                                    placeholder="Enter your email"
                                                    required
                                                />
                                            </div>
                                            <div className="mb-3">
                                                <label
                                                    for="exampleInputPassword1"
                                                    className="form-label login-page-right-form-label"
                                                >
                                                    Password*
                                                </label>
                                                <input
                                                    type="password"
                                                    className="form-control login-page-right-form-input"
                                                    id="exampleInputPassword1"
                                                    placeholder="Create a password"
                                                    required
                                                />
                                            </div>
                                            <div className="mb-3 signup-page-right-form-password-info">
                                                <p className="text-start">
                                                    Must be at least 8
                                                    characters.{" "}
                                                </p>
                                            </div>
                                            <button
                                                type="submit"
                                                className=" my-4 btn btn-primary w-100 sign-in-button"
                                            >
                                                Get started
                                            </button>
                                        </form>

                                        <button className="btn btn-outline-dark w-100 sign-in-button sign-in-with-google">
                                            <img src="../images/google-icon.png"></img>
                                            <span className="text-dark m-0 px-3 py-2">
                                                Sign up with Google
                                            </span>
                                        </button>
                                    </div>
                                </div>
                            </div>
                            <div className="col-lg-12 mt-3 dont-have-a-account">
                                <p className="text-secondary text-center m-0 px-3 py-2">
                                    Already have an account?
                                    <a className="ms-2 text-decoration-none">
                                        Log in
                                    </a>
                                </p>
                            </div>
                        </div>
                        <div className="col-lg-2"></div>
                    </div>
                </div>
            </div>
        </>
    );
}
