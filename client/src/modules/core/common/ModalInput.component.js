import React from "react";

export default function ModalInput({ label, inputId, placeholder }) {
    return (
        <>
            <div className="mb-3 row ">
                <label
                    htmlFor={inputId}
                    className="col-sm-4 col-form-label modal-input-label"
                >
                    {label}
                </label>
                <div className="col-sm-8 modal-input">
                    <input
                        type="text"
                        readOnly
                        className="form-control"
                        id={inputId}
                        placeholder={placeholder}
                    />
                </div>
            </div>
        </>
    );
}
