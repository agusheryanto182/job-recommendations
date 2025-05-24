interface LinkProps {
  href: string;
  children: React.ReactNode;
  className?: string;
}

export function Link({ href, children, className = '' }: LinkProps) {
  return (
    <a
      href={href}
      className={`text-foreground/70 hover:text-foreground transition-colors ${className}`}
    >
      {children}
    </a>
  );
}
