import { Navbar } from '@/components/organisms/navbar';
import { Footer } from '@/components/organisms/footer';
import { Hero } from '@/components/organisms/hero';
export function LandingPageTemplate() {
  return (
    <main>
      <Navbar />
      <Hero />
      <Footer />
    </main>
  );
}
