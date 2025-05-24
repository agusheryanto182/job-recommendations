"use client";
import React from "react";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { AuroraText } from "@/components/magicui/aurora-text";
import { Container } from "@/components/atoms/container";

export function Hero() {
  return (
    <Container className="relative min-h-screen mt-16">
      <div className="absolute top-[-20%] left-[50%] -translate-x-1/2 w-[1200px] h-[400px] rounded-full bg-gradient-to-r from-blue-500/30 to-purple-500/30 blur-[120px] pointer-events-none animate-pulse" />
      <div className="absolute top-[10%] right-[-10%] w-[600px] h-[600px] rounded-full bg-gradient-to-bl from-purple-600/25 to-blue-400/25 blur-[130px] pointer-events-none animate-pulse" />
      <div className="absolute bottom-[5%] left-[-5%] w-[500px] h-[500px] rounded-full bg-gradient-to-tr from-blue-500/20 to-purple-500/20 blur-[100px] pointer-events-none animate-pulse" />

      <div className="absolute top-[30%] left-[20%] w-[300px] h-[300px] rounded-full bg-blue-300/10 blur-[80px] pointer-events-none" />
      <div className="absolute bottom-[20%] right-[10%] w-[400px] h-[400px] rounded-full bg-purple-300/10 blur-[90px] pointer-events-none" />

      <div className="absolute top-[40%] left-[30%] w-[200px] h-[200px] rounded-full bg-gradient-to-r from-blue-400/20 to-purple-400/20 blur-[50px] pointer-events-none animate-pulse" />
      <div className="absolute bottom-[30%] right-[25%] w-[250px] h-[250px] rounded-full bg-gradient-to-l from-purple-400/15 to-blue-400/15 blur-[60px] pointer-events-none animate-pulse" />

      <div className="container mx-auto px-6 py-20 flex flex-col lg:flex-row items-center justify-between gap-12 relative">
        <div className="flex flex-col gap-6 max-w-xl relative">
          <div className="absolute -left-4 -top-4 w-[120%] h-[120%] bg-gradient-to-r from-blue-500/5 to-purple-500/5 blur-2xl rounded-full pointer-events-none" />

          <h1 className="text-4xl sm:text-6xl font-extrabold relative">
            <span className="text-black">Job Matching</span>
            <br />
            <span className="bg-gradient-to-r from-blue-500 to-purple-600 bg-clip-text text-transparent">
              Powered by AI
            </span>
          </h1>

          <p className="text-gray-700 text-lg leading-relaxed relative">
            Biarkan AI kami menganalisis CV-mu dan menemukan pekerjaan yang
            sempurna untukmu dari{" "}
            <span className="font-semibold">ribuan lowongan teknologi</span>{" "}
            dalam hitungan detik.
          </p>

          <div className="flex flex-wrap gap-4 pt-4">
            <Card className="border border-white/20 rounded-xl bg-white/70 backdrop-blur-xl px-6 py-4 text-center text-black hover:bg-white/80 transition-all duration-300">
              <p className="text-lg font-bold">
                <AuroraText>10.000+</AuroraText>
              </p>
              <p className="text-sm text-gray-500 font-bold">MATCHED TALENTS</p>
            </Card>
            <Card className="border border-white/20 rounded-xl bg-white/70 backdrop-blur-xl px-6 py-4 text-center text-black hover:bg-white/80 transition-all duration-300">
              <p className="text-lg font-bold">
                <AuroraText>95%</AuroraText>
              </p>
              <p className="text-sm text-gray-500 font-bold">MATCH ACCURACY</p>
            </Card>
            <Card className="border border-white/20 rounded-xl bg-white/70 backdrop-blur-xl px-6 py-4 text-center text-black hover:bg-white/80 transition-all duration-300">
              <p className="text-lg font-bold">
                <AuroraText>5000+</AuroraText>
              </p>
              <p className="text-sm text-gray-500 font-bold">TECH COMPANIES</p>
            </Card>
          </div>
        </div>

        <Card className="relative bg-white/70 backdrop-blur-xl rounded-4xl px-10 py-20 w-full max-w-lg h-[500px] text-center flex flex-col items-center justify-center gap-6 border border-white/20">
          <div className="absolute inset-0 bg-gradient-to-br from-blue-500/5 to-purple-500/5 rounded-4xl blur-xl" />

          <h2 className="text-2xl font-bold mb-2 relative">
            Temukan Match Karirmu
          </h2>
          <p className="text-gray-600 mb-6 relative">
            Upload CV dan dapatkan rekomendasi dalam hitungan detik
          </p>
          <Button className="w-full bg-gradient-to-br from-blue-500 to-purple-600 text-white hover:opacity-90 transition-all duration-300 relative hover:scale-[1.02]">
            Masuk dengan Google
          </Button>
        </Card>
      </div>
    </Container>
  );
}
