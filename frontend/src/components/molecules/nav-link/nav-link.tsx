interface NavLinkProps {
  href: string;
  children: React.ReactNode;
  active?: boolean;
}

export function NavLink({ href, children, active }: NavLinkProps) {
  return (
    <a
      href={href}
      className={`
        relative px-4 py-2 rounded-full
        font-extrabold text-sm
        ${active ? "text-blue-500" : "text-gray-500"}
        hover:bg-background/10
        group
      `}
    >
      <span
        className={`
        absolute inset-0 rounded-full opacity-0 
        bg-gradient-to-tr from-white/5 to-white/[0.02]
        group-hover:opacity-100
      `}
      />

      {active ? (
        <span className="relative z-10 bg-gradient-to-r from-blue-500 to-purple-600 bg-clip-text text-transparent">
          {children}
        </span>
      ) : (
        <span className="relative z-10">{children}</span>
      )}

      {active && (
        <span className="absolute bottom-0 left-0 w-full h-[3px] bg-gradient-to-r from-blue-500 to-purple-600" />
      )}
    </a>
  );
}
