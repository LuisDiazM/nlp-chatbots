import { useFormik } from "formik";

interface youtubeFormData {
  name: string;
  email: string;
  channel: string;
}

const YouTubeForm = () => {
  const initialValues: youtubeFormData = {
    name: "",
    email: "",
    channel: "",
  };
  const formik = useFormik({
    initialValues,
    onSubmit: () => {
      console.log(formik.values);
    },
  });
  return (
    <>
      <form onSubmit={formik.handleSubmit}>
        <label htmlFor="name">Nombre</label>
        <input
          type="text"
          onChange={formik.handleChange}
          value={formik.values.name}
          name="name"
        ></input>

        <label htmlFor="email">Email</label>
        <input
          type="email"
          onChange={formik.handleChange}
          value={formik.values.email}
          name="email"
        ></input>

        <label htmlFor="channel">Canal</label>
        <input
          type="text"
          onChange={formik.handleChange}
          value={formik.values.channel}
          name="channel"
        ></input>
        <button type="submit">Enviar</button>
      </form>
    </>
  );
};

export default YouTubeForm;
