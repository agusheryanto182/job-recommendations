"use client";
import { motion } from "motion/react";
import { Target } from "lucide-react";

export function Logo() {
  return (
    <motion.div
      className="flex items-center gap-3 cursor-pointer"
      whileHover={{ scale: 1.02 }}
      whileTap={{ scale: 0.98 }}
    >
      <motion.div
        className="relative flex items-center justify-center p-3 rounded-xl shadow-lg"
        animate={{
          background: "linear-gradient(45deg, #3b82f6 0%, #8b5cf6 100%)",
          transition: {
            duration: 3,
            repeat: Infinity,
            repeatType: "reverse",
          },
        }}
      >
        <Target className="w-5 h-5 text-white" />
      </motion.div>

      <div className="flex flex-col">
        <motion.div
          className="text-2xl font-extrabold tracking-tight"
          initial={{ opacity: 0, y: 5 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
        >
          <span className="text-gray-800">Job</span>
          <span className="bg-clip-text text-transparent bg-gradient-to-r from-blue-500 to-purple-600">
            Matcher
          </span>
        </motion.div>
        <motion.div
          className="text-xs text-gray-500 font-medium -mt-1"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 0.3, duration: 0.5 }}
        >
          AI-Powered Job Matching
        </motion.div>
      </div>
    </motion.div>
  );
}
