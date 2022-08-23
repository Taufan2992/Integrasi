import React, { useState, useEffect, useContext } from 'react';
import { UserContext } from '../../context/user-context';
import { Container, Row, Col } from "react-bootstrap"
import Rp from "rupiah-format"
import "../../assets/css/DetailProduct.css"
import { useParams, useNavigate } from 'react-router-dom';
import NavbarUser from "../../components/partials/NavbarUser";
import { useQuery, useMutation } from "react-query"
import { API } from "../../config/API"

const DetailProduct = () => {
    const title = ' Detail Product '
    document.title = title 
    const [cartCounter, setCartCounter] = useState(0)

    localStorage.setItem("Tambah", cartCounter)
    const addCart = localStorage.getItem("Tambah")

const moving = useNavigate()

const params = useParams();
const id = params.id;

// toping
const [topping, setTopping] = useState([]);
const [topping_id, setTopping_id] = useState ([]);

const handleChange = (e) => {
  let updateTopping = [...topping];
  if (e.target.checked) {
    updateTopping = [...topping, e.target.value];
  } else {
    updateTopping.splice(topping.indexOf(e.target.name));
  }
  setTopping(updateTopping);

let toppingId = [...topping_id];
  if (e.target.checked) {
    toppingId = [...topping_id, parseInt(e.target.name)];
  } else {
    toppingId.splice(topping_id.indexOf(e.target.name));
  }

  setTopping_id(toppingId);
};

let { data: product } = useQuery('productCache', async () => {
  const res = await API.get(`/product/${id}`)
  return res.data.data
})

let { data: toppings } = useQuery('topingsCache', async () => {
  const response = await API.get('/topings');
  return response.data.data;
});

// tambah price
let ToppingTotal = topping.reduce((a, b) => {
  return a + parseInt(b);
}, 0);

let subamount = product?.price + ToppingTotal;
let qty = 1;

const handleSubmit = useMutation(async (e) => {
  try {
    e.preventDefault();

    const config = {
      headers: {
        "Content-type": "application/json",
      },
    };
    const body = JSON.stringify({
      toping_id: topping_id,
      subamount: subamount,
      product_id: parseInt(id),
      qty: qty
    });

    const response = await API.post("/cart", body, config);
    console.log(response);
    setCartCounter(cartCounter + 1)
    alert('Data added succesfully')
    if(response.status == 200){
      moving('/cart')
    }

    // navigate("/");
  } catch (error) {
    console.log(error);
  }
});

    return (
        <Container>
            <NavbarUser plusOne={addCart}/>
            <Row id="row-detail-product">
                <Col className="detail-drink mt-5">
                    <img id="detail-img-drink" className='mt-5 mb-5 shadow-lg' style={{objectFit: "cover"}} src={product?.image}/>
                </Col>
                <Col id="right-side-addtpg" className="mt-5">
                    <div className="title-detail-product">
                        <p className="mt-4">{product?.title}</p>
                    </div>
                    <div className="price-drink">
                        <p className="mt-2">{Rp.convert(product?.price)}</p>
                    </div>
                    <div className="toping-add">
                        <p className='mt-5'>Toping</p>
                    </div>

                    {/* MAPPING TOPPING */}
                    <Row>
                    {toppings?.map((data, index) => (
                        <div key={index} className="topping-datas ms-4 col mb-5">
                            <div className="img-data-toping toppings-list-item" >
                                <div>
                                    <input 
                                        type="checkbox" 
                                        className="poppingCheck" 
                                        style={{display:"none"}}
                                        value={data.price}
                                        name={data.id} 
                                        id={`checkmark${index}`} 
                                        onChange={handleChange} 
                                    />
                                    <label htmlFor={`checkmark${index}`}>
                                        <img 
                                          className="mb-5 cursor-pointer" 
                                          style={{objectFit: "cover"}}  
                                          width="200px" src={data.image} 
                                        />
                                    </label>
                                    
                                    <p id="toping-name" className="mb-5">{data.title}</p>
                                </div>
                            </div>
                            <div className="price-data-toping ms-4 mb-5" hidden>
                                <p>{data?.price}</p>
                            </div>
                        </div>
                    ))}
                    
                    {/* END MAPPING */}

                        <div className="sub-total d-flex mb-5">
                            <div className="left-total">
                                Total
                            </div>
                            <div className="right-total">
                                {Rp.convert(product?.price + ToppingTotal)}
                            </div>
                        </div>
                        <div className="btn-add-cart mb-5 mt-2">
                            <button className="mb-2" onClick={(e) => handleSubmit.mutate(e)}>Add Cart</button>
                        </div>
                    </Row>
                </Col>
            </Row>
        </Container>
    )
}

export default DetailProduct