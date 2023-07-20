// import { getGalleryById } from "@/services/gallery";
import Customer from "@/components/customer/Customer";
import { NextPage } from "next";
import Image from "next/image";

interface Props {

}

// const getPost = async (params: Props["params"]) => {
//   try {
//     const res = await axios.get(
//       "https://api.slingacademy.com/v1/sample-data/photos/" + params.id
//     );

//     return res.data;
//   } catch (error) {}
// };

const page: NextPage<Props> = async ({  }) => {
  // const post = await getGalleryById(params.id);

  return (
    <div className="pt-24">
      <Customer/>
    </div>
  );
};

export default page;
