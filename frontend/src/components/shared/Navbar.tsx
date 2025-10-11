import Link from "next/link";
import LionHeartLogo from "./LionHeartLogo";

export default function Navbar() {
  const sections = [
    {
      name: "Campaigns",
      link: "#campaigns",
    },
    {
      name: "Events",
      link: "#events",
    },
    {
      name: "About Us",
      link: "#aboutus",
    },
    {
      name: "Contact",
      link: "#contact",
    },
  ];

  return (
    <nav className="sticky">
      <div className="flex items-center justify-between w-[calc(100%-2*clamp(2rem,1.43rem+2.86vw,4rem))] mx-auto max-w-[90rem]">
        <div className="w-18 h-18 py-4 flex-shrink-0">
          <Link href="/">
            <LionHeartLogo
              classname="h-full w-full"
            />
          </Link>
        </div>
        <div className="flex gap-8 font-medium text-lg text-text-300 font-work-sans">
          {sections.map((item, index) => (
            <div key={index}>
              <a href={item.link} className="hover:text-text-100">
                {item.name}
              </a>
            </div>
          ))}
        </div>
      </div>
    </nav>
  );
}
