import React, { useState, useEffect } from 'react';
import Navbar from "../../components/partials/NavbarAdmin"
import { Container, Row, Col, Card, Button } from "react-bootstrap"
import { useQuery, useMutation } from "react-query"
import { API } from "../../config/API"
import "../../assets/css/Auth.css"
import Rp from "rupiah-format"
import Delete from "../../components/modal/DeleteData.jsx"
import { useNavigate } from 'react-router-dom';

const MainAdmin = () => {
    const ttl = "Home Admin"
    document.title = ttl
    
    const [deleteOne, setDeleteOne] = useState(null)
    const [confirmDelete,setConfirmDelete] = useState(null)

    const [show, setShow] = useState(false)
    const handleClose = () => setShow(false)
    const handleShow = () => setShow(true)

    const moving = useNavigate()
        //   GET PRODUCTS
        let {data: products, refetch} = useQuery('productsCache', async () => {
            const response = await API.get('/products')
            return response.data.data
              })
              console.log('====================================');
              console.log(products);
              console.log('====================================');
        // DELETE 
        const handleDelete = (id) => {
            setDeleteOne(id);
            handleShow();
          };

        //   MUTATE DELETE
    const deleteById = useMutation(async (id) => {
        try {
          const response = await API.delete(`/product/${id}`);
          console.log(response);

          refetch()
        } catch (error) {
          console.log(error);
        }
      });

          // UPDATE PRODUCT
    const editProduct = (id) => {
        moving('/update-product/' + id)
    }

          //  LIFECYCLE DELETE
    useEffect(() => {
        if (confirmDelete) {
          // Close modal confirm delete data
          handleClose();
          // execute delete data by id function
          deleteById.mutate(deleteOne);
          setConfirmDelete(null);
        }
      }, [confirmDelete]);

    return (
        <>
              {/* LIST PRODUCTS */}
      <Container className="ms-5" >
        <Navbar/>
        <Row className="ms-5">
            <div className="footer-title mt-5">
                <p className="ms-5 mb-5">
                    List Products
                </p>
            </div>
            {products?.map((item, index) => (
                <Col className="mapping-card ms-5 mb-5" key={index}>
                    <Card id="card-main-admin">
                        <div className="img-drink">
                            <Card.Img id="per-img-product" variant="top" src={item?.image}/>
                        </div>
                        <div className="name-drink ms-2 mt-3">
                            <p>{item?.title}</p>
                        </div>
                        <div className="price-drink ms-2">
                            <p>{Rp.convert(item?.price)}</p>
                        </div>
                        <div className="btn-edit-delete mb-2">
                            <Button onClick={() => {
                                editProduct(item.id);
                            }}className="btn-list-products ms-2 me-2" variant="primary">Edit</Button>
                            <Button onClick={() => {
                                handleDelete(item.id);
                            }}className="btn-list-products ms-2" variant="danger">Delete</Button>
                        </div>
                    </Card>
                </Col>
            ))}
            <Delete
                setConfirmDelete={setConfirmDelete}
                show={show}
                handleClose={handleClose}
            />
        </Row>
    </Container>
        {/* END LIST PRODUCTS */}
</>
    )
}

export default MainAdmin