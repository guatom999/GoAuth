import axios from "axios"

interface IProductsPrice {
        limit? : number,
}


export  const getProductsPrice = async (arg : IProductsPrice) => {
    
    try {
        const { data } = await axios.get("http://127.0.0.1:8000/totalproductsprice")

        return data
    }catch(error){
        console.log(error)
        return error
    }

}