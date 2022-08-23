import React, { useState, UserContext } from 'react';
import Bin from "../../assets/img/bin.png";
import NavbarUser from "../../components/partials/NavbarUser";
import "../../assets/css/Cart.css"
import Rp from "rupiah-format"
import { Alert , Form } from "react-bootstrap"
import { useMutation, useQuery } from "react-query";
import { API } from "../../config/API"
import {  useEffect } from "react";
import { useNavigate } from "react-router-dom";


function Cart() {
  
  const title = "My Cart"
  document.title = title

  const [message, setMessage] = useState(null)
  // const [state, dispatch] = useContext(UserContext);
  // modal
  // const [showTrans, setShowTrans] = useState(false);

  // const handleClose = () => setShowTrans(false);
  let navigate = useNavigate();

  // cart
  let { data: cart, refetch } = useQuery("cartsCache", async () => {
    const response = await API.get("/carts-id");
    return response.data.data;
  });

  // subtotal
  let resultTotal = cart?.reduce((a, b) => {
    return a + b.subamount;
  }, 0);

  // remove
  let handleDelete = async (idx) => {
    console.log(idx);
    await API.delete(`/cart/${idx}`);
    refetch();
  };

    // pay
    const form = {
      status: "failed",
      amount: resultTotal,
    };
    const handleSubmit = useMutation(async (e) => {
      e.preventDefault();

      const alert = (
        <Alert id="alert-message-payment" variant="success" className='py-3'>
          Thank you for ordering in us, please wait to verify you order
        </Alert>
      )
      setMessage(alert)

      const config = {
        headers: {
          "Content-type": "application/json",
        },
      };
  
      const body = JSON.stringify(form);
  
      const response = await API.patch("/transactionid", body, config);
      console.log(response);
      const token = response.data.data.token;
      console.log(token);

    window.snap.pay(token, {
      onSuccess: function (result) {
        /* You may add your own implementation here */
        console.log(result);
        navigate("/profile");
      },
      onPending: function (result) {
        /* You may add your own implementation here */
        console.log(result);
        navigate("/profile");
      },
      onError: function (result) {
        /* You may add your own implementation here */
        console.log(result);
      },
      onClose: function () {
        /* You may add your own implementation here */
        alert("you closed the popup without finishing the payment");
      },
    });
    });

    useEffect(() => {
      //change this to the script source you want to load, for example this is snap.js sandbox env
      const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
      //change this according to your client-key
      const myMidtransClientKey = "Client key here ...";
  
      let scriptTag = document.createElement("script");
      scriptTag.src = midtransScriptUrl;
      // optional if you want to set script attribute
      // for example snap.js have data-client-key attribute
      scriptTag.setAttribute("data-client-key", myMidtransClientKey);
  
      document.body.appendChild(scriptTag);
      return () => {
        document.body.removeChild(scriptTag);
      };
    }, []);

  let addCart = localStorage.getItem("Tambah")
  return (
    <div>
      <div className="container">
        <NavbarUser className="ms-4 for-nav" plusOne={addCart}/>
        {message}
      </div>
      <div className="p-5 mx-5">
        <div className="px-5 mb-3 text-red">
          <h3>My Cart</h3>
        </div>
        <div className="px-5">
          <p  className="mb-0 text-red">Review your order</p>
        </div>

        <div className="row">
          <div className="col-8 px-5">
            <hr />

            <div className="card mb-3 scroll" style={{ border: "none" }}>
              {cart?.map((item, index) => (
                <div className="row g-0 mb-2" key={index}>
                  <div className="col-md-2">
                    <img
                      src={item?.product?.image}
                      alt="cartimage"
                      className="rounded"
                      height={"100px"}
                      width={"100px"}
                      style={{ objectFit: "cover" }}
                    />
                  </div>
                  <div className="col-md-10 row">
                    <div className="col-md-9">
                      <div className="card-body px-0">
                        <p
                          className="card-title text-red"
                          style={{ fontSize: "18px", fontWeight: "900" }}
                        >
                          {item.product.title}
                        </p>

                        <p
                          className="card-text"
                          style={{ fontSize: "16px", fontWeight: "800", color:"#974A4A" }}
                        >
                          Toping:{" "}
                          <span 
                          className="text-red ms-1"
                          style={{fontSize:"14px", fontWeight: "100"}}>
                            {" "}
                            {item.toping?.map((item, idx) => (
                              <span className="d-inline" key={idx}>
                                {item.title},
                              </span>
                            ))}
                          </span>
                        </p>
                      </div>
                    </div>

                    <div className="col-md-3">
                      <div className="card-body px-0">
                        <p
                          className="card-title text-red"
                          style={{
                            fontSize: "16px",
                            fontWeight: "400",
                            textAlign: "right",
                          }}
                        >
                          {Rp.convert(item?.subamount)}
                        </p>

                        <img src={Bin} style={{ float: "right" , cursor:'pointer'}}
                           onClick={() => handleDelete(item?.id)}/>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>

            <hr />


          </div>
            <div className="col-4 px-5">

                <div className="text-red">
                  <hr />
                  <div className="d-flex justify-content-between">
                    <p className="d-flex">Subtotal</p>
                    <p className="d-flex">{Rp.convert(resultTotal)}</p>
                  </div>
                  <div className="d-flex justify-content-between">
                    <p className="d-flex">Qty</p>
                    <p className="d-flex">{cart?.length}</p>
                  </div>
                  <hr />
                  <div className="d-flex justify-content-between">
                    <p className="d-flex fw-bold">Total</p>
                    <p className="d-flex fw-bold">{Rp.convert(resultTotal)}</p>
                  </div>
                </div>
              <div className="mt-4">
              <Form className="d-flex">
                <button 
                className="container btn btn-primary bg-red border-0 mt-2" 
                type="submit"
                onClick={(e) => handleSubmit.mutate(e)}
                >
                  Pay
                </button>
              </Form>
              </div>
            </div>
        </div>
      </div>
    </div>
  );
}

export default Cart;
