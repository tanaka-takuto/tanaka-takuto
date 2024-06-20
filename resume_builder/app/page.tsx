import Image from "next/image";

export default function Home() {
  return (
    <>
      <Image
        src={"/resume_builder.webp"}
        width={500}
        height={500}
        alt="resume_builder"
      />
      <button className="btn btn-primary">Hello, DaisyUI</button>
    </>
  );
}
