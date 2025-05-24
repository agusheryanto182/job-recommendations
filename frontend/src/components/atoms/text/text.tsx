import { JSX } from 'react'

type TextVariant = 'heading1' | 'heading2' | 'body' | 'small'

interface TextProps {
  variant?: TextVariant
  children: React.ReactNode
  className?: string
}

const TEXT_STYLES: Record<TextVariant, string> = {
  heading1: 'text-4xl font-bold',
  heading2: 'text-2xl font-semibold',
  body: 'text-base',
  small: 'text-sm text-muted-foreground',
} as const

const TAG_MAPPING: Record<TextVariant, keyof JSX.IntrinsicElements> = {
  heading1: 'h1',
  heading2: 'h2',
  body: 'p',
  small: 'small',
} as const

/**
 * Text component for consistent typography across the application
 * @param props.variant - The style variant of the text (heading1 | heading2 | body | small)
 * @param props.children - The content to be rendered
 * @param props.className - Additional CSS classes to be applied
 */
export function Text({
  variant = 'body',
  children,
  className = '',
}: TextProps) {
  const Tag = TAG_MAPPING[variant]

  return (
    <Tag className={`${TEXT_STYLES[variant]} ${className}`}>{children}</Tag>
  )
}
