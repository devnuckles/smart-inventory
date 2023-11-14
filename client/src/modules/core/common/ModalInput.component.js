import React from "react";

export default function ModalInput({
  label,
  inputId,
  placeholder,
  name,
  value,
  onChange,
}) {
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
            className="form-control"
            id={inputId}
            placeholder={placeholder}
            name={name}
            onChange={onChange}
          />
        </div>
      </div>
    </>
  );
}
