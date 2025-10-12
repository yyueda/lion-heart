import { cn } from "@/lib/utils";

type HeaderProps = {
  title: string
  desc?: string
  className?: string
}

export default function Header({
  title,
  desc = "",
  className = ""
}: HeaderProps) {
  return (
    <header 
      className={cn("flex flex-col gap-2 bg-bg-100 border border-border-100 rounded-md p-4 text-text-100 w-full", className)}
    >
      <h1
        className="text-2xl font-bold"
      >
        {title}
      </h1>
      <p
        className="text-text-300"
      >
        {desc}
      </p>
    </header>
  );
}
