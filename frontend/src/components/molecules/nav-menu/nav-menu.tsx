import { NavLink } from "@/components/molecules/nav-link";
import { useActiveSection } from "@/hooks/use-active-section/use-active-section";

export function NavMenu() {
  const activeSection = useActiveSection();

  return (
    <div className="flex md:flex-row flex-col items-center gap-8 text-sm  whitespace-nowrap">
      <NavLink href="#home" active={activeSection === "home"}>
        HOME
      </NavLink>
      <NavLink href="#hasil" active={activeSection === "hasil"}>
        HASIL
      </NavLink>
      <NavLink href="#paket" active={activeSection === "paket"}>
        PAKET
      </NavLink>
    </div>
  );
}
