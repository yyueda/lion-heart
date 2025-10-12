import AppSidebar from "@/components/shared/AppSidebar";
import { SidebarProvider } from "@/components/ui/sidebar";

export default function EventsLayout({
    children,
}: Readonly<{
  children: React.ReactNode;
}>) {
    return (
        <SidebarProvider className="gap-4">
            <AppSidebar />
            <main>
                {children}
            </main>
        </SidebarProvider>
    );
}
