import React, { useState, useEffect, useContext } from 'react';
import { Container, Row, Col, Form } from "react-bootstrap"
import "../../assets/css/AddProduct.css"
import Navbar from "../../components/partials/NavbarUser"
import { useQuery, useMutation } from 'react-query';
import { API } from '../../config/API';
import { UserContext } from '../../context/user-context';
import IconUpload from "../../assets/img/ikon-upload.png"
import NoImg from "../../assets/img/blank-profile.png"
import { useNavigate, useParams } from 'react-router-dom';


const EditProfile = () => {

    const ttl = "Edit Profile"
    document.title = ttl

    const [state, _] = useContext(UserContext)

    const [preview, setPreview] = useState(null)
    const { id } = state.user

    const moving = useNavigate()
    console.log(state.user.id);
    const [fetchProfile, setFetchProfile] = useState({
        name : "",
        email : "",
        image : ""
    })


    // Fetch data
    const { data : profileData } = useQuery('profileCache', async () => {
        const res = await API.get('/user/' + id)
        return res.data.data
    })

    useEffect(() => {
        setPreview(profileData?.image)
        setFetchProfile({
            ...profileData,
            name : profileData?.name,
            email : profileData?.email
        })
    },[profileData])

    const handleOnChange = (e) => {
        setFetchProfile({
            ...fetchProfile,
            [e.target.name]: e.target.type === "file" ? e.target.files : e.target.value
        })
        if(e.target.type === 'file'){
            let url = URL.createObjectURL(e.target.files[0])
            setPreview(url)
        }
    }
    
    const handleOnSubmit = useMutation(async (e) => {
        try {
            e.preventDefault()

            // Configuration
            const config = {
                method: "PATCH",
                headers: {
                  'Content-type': 'multipart/form-data',
                },
              };

            const data = new FormData()
            
            if (fetchProfile?.image) {
                data.set("image", fetchProfile?.image[0],fetchProfile?.image[0].name);
              }
            data.set("name", fetchProfile?.name)
            data.set("email", fetchProfile?.email)
            console.log(fetchProfile?.image);
            const response = await API.patch(`/user/${id}`, data, config);
            console.log(response)
            alert('Success update!')
            moving("/profile")
        } catch (error) {
            console.log(error);
        }
    })
    console.log(preview);
    return (
        <Container>
        <Navbar/>
        <Row className="ms-5">
            <Col id="left-side-form" className="mt-4">
                <div className="header-title mt-5">
                    <p className="title-edit-profile mb-5">
                        Update Profile
                    </p>
                </div>
                <Form onSubmit={(e) => handleOnSubmit.mutate(e)} >
                    <Form.Group className="mb-4" controlId="formInputProduct">
                        <Form.Control name="name" autoComplete="off" className="formInputProduct" type="text" placeholder="Your Name" onChange={handleOnChange}
                        value={fetchProfile?.name}
                        />
                    </Form.Group>
                    <Form.Group className="mb-2 mt-4" controlId="formInputProduct">
                        <Form.Control name="email" autoComplete="off" className="formInputProduct mt-4" 
                         type="text" placeholder="Your Email" onChange={handleOnChange}
                         value={fetchProfile?.email}
                        />
                    </Form.Group>
                    <Form.Group className="mb-4" controlId="formInputProduct">
                        <input
                        type="file"
                        id="upload"
                        name="image"
                        onChange={handleOnChange}
                        hidden
                        />
                        <label for="upload" className="label-file-add-product">
                            <img className="position-absolute" src={IconUpload}/>
                        </label>
                        <Form.Control name="image" className="formInputProduct" type="text" placeholder="Your Photo"
                         onChange={handleOnChange} value={preview}/>
                    </Form.Group>
                    <div className="btn-submit-prdct ms-5">
                        <button type='submit'>Edit Profile</button>
                    </div>
                </Form>
            </Col>
            <Col className="ms-4 mt-5">
                <div className="img-detail-product ms-3 mt-3 mb-5">
                    <img src={ preview || NoImg } />
                </div>
            </Col>
        </Row>
    </Container>

    )
}

export default EditProfile