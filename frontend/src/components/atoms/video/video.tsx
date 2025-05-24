'use client';
import React, { useState } from 'react';
import dynamic from 'next/dynamic';
import { Play } from 'lucide-react';
import Image from 'next/image';
import { Container } from '@/components/atoms/container';
const ReactPlayer = dynamic(() => import('react-player'), {
  ssr: false,
  loading: () => (
    <div className="w-full h-full bg-gray-100 animate-pulse rounded-xl" />
  ),
});

interface VideoProps {
  src: string;
  thumbnail?: string;
  className?: string;
}

export function Video({ src, thumbnail, className = '' }: VideoProps) {
  const [showModal, setShowModal] = useState(false);

  return (
    <>
      <Container>
        <div
          className={`relative aspect-video rounded-xl overflow-hidden cursor-pointer group ${className}`}
          onClick={() => setShowModal(true)}
          aria-label="Play video"
        >
          <div className="w-full h-full bg-gray-900">
            {thumbnail ? (
              <Image
                src={thumbnail}
                alt="Video thumbnail"
                className="w-full h-full object-cover"
                width={1920}
                height={1080}
                quality={100}
                priority
              />
            ) : (
              <div className="w-full h-full bg-gray-900" />
            )}
          </div>

          <div className="absolute inset-0 flex items-center justify-center">
            <div className="rounded-full bg-blue-600/90 p-4 transform transition-all duration-200 md:group-hover:scale-150 group-hover:scale-125 group-hover:bg-blue-600">
              <Play className="md:w-8 md:h-8 w-6 h-6 text-white" />
            </div>
          </div>
        </div>
      </Container>

      {showModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-sm">
          <button
            onClick={() => setShowModal(false)}
            className="absolute top-4 right-4 text-white hover:text-gray-300 p-2"
            aria-label="Close video"
          >
            <svg
              className="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>

          <div className="w-full max-w-6xl mx-4 aspect-video">
            <ReactPlayer
              url={src}
              width="100%"
              height="100%"
              controls={true}
              playing={true}
              config={{
                youtube: {
                  playerVars: {
                    showinfo: 1,
                    origin:
                      typeof window !== 'undefined'
                        ? window.location.origin
                        : '',
                  },
                },
              }}
            />
          </div>
        </div>
      )}
    </>
  );
}
