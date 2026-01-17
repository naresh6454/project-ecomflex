import type { Config } from 'tailwindcss'
import plugin from 'tailwindcss/plugin'

export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        // Primary colors
        'accent': {
          DEFAULT: '#FFD700', // Gold accent
          'light': '#FFEC80',
          'dark': '#E6C200', // Darker gold for hover states
        },
        
        // Brand colors - blue theme
        'brand': {
          DEFAULT: '#3B82F6', // Main blue (alias for 500)
          'light': '#EFF6FF', // Light blue background (alias for 50)
          'dark': '#2563EB', // Deep blue (alias for 600)
          50: '#EFF6FF',  // Lightest blue
          100: '#DBEAFE', // Very light blue
          200: '#BFDBFE', // Light blue
          300: '#93C5FD', // Medium light blue
          400: '#60A5FA', // Medium blue
          500: '#3B82F6', // Main blue
          600: '#2563EB', // Medium dark blue
          700: '#1D4ED8', // Dark blue
          800: '#1E40AF', // Very dark blue
          900: '#1E3A8A', // Darkest blue
          950: '#172554', // Nearly black blue
        },
        
        // Neutral colors
        'neutral': {
          'bg': '#FFFFFF',      // Card backgrounds
          'text': '#1F2937',    // Main text
          'border': '#E5E7EB',  // Borders
          'muted': '#9CA3AF',   // Subdued text
        },
        
        // Status colors
        'success': {
          DEFAULT: '#22C55E',
          'light': '#DCFCE7',
        },
        'warning': {
          DEFAULT: '#F59E0B',
          'light': '#FEF3C7',
        },
        'error': {
          DEFAULT: '#EF4444',
          'light': '#FEE2E2',
        },
        'info': {
          DEFAULT: '#3B82F6',
          'light': '#DBEAFE',
        },
      },
      fontFamily: {
        'sans': ['Inter', 'system-ui', 'sans-serif'],
        'roboto': ['Roboto', 'sans-serif'],
      },
      fontSize: {
        'xs': ['0.75rem', { lineHeight: '1rem' }],
        'sm': ['0.875rem', { lineHeight: '1.25rem' }],
        'base': ['1rem', { lineHeight: '1.5rem' }],
        'lg': ['1.125rem', { lineHeight: '1.75rem' }],
        'xl': ['1.25rem', { lineHeight: '1.75rem' }],
        '2xl': ['1.5rem', { lineHeight: '2rem' }],
        '3xl': ['1.875rem', { lineHeight: '2.25rem' }],
        '4xl': ['2.25rem', { lineHeight: '2.5rem' }],
        '5xl': ['3rem', { lineHeight: '1' }],
        '6xl': ['3.75rem', { lineHeight: '1' }],
      },
      boxShadow: {
        'card': '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
        'card-hover': '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
        'button': '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)',
        'golden': '0 0 15px rgba(255, 215, 0, 0.5)',
      },
      borderRadius: {
        'xl': '0.75rem',
        '2xl': '1rem',
        '3xl': '1.5rem',
      },
      animation: {
        'fade-in': 'fadeIn 0.5s ease-in-out',
        'slide-up': 'slideUp 0.5s ease-out',
        'slide-down': 'slideDown 0.5s ease-out',
        'slide-left': 'slideLeft 0.3s ease-out',
        'slide-right': 'slideRight 0.3s ease-out',
        'scale-in': 'scaleIn 0.3s ease-out',
        'pulse-gold': 'pulseGold 2s infinite',
        'bounce-subtle': 'bounceSlight 1.5s infinite',
        'shimmer': 'shimmer 2.5s infinite',
        'pulse': 'pulse 2s infinite',
        'golden-pulse': 'goldenPulse 1.5s infinite',
        'spin-slow': 'spin 3s linear infinite',
        'float': 'float 3s ease-in-out infinite',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
        slideUp: {
          '0%': { transform: 'translateY(20px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' },
        },
        slideDown: {
          '0%': { transform: 'translateY(-20px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' },
        },
        slideLeft: {
          '0%': { opacity: '0', transform: 'translateX(20px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' },
        },
        slideRight: {
          '0%': { opacity: '0', transform: 'translateX(-20px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' },
        },
        scaleIn: {
          '0%': { transform: 'scale(0.95)', opacity: '0' },
          '100%': { transform: 'scale(1)', opacity: '1' },
        },
        pulseGold: {
          '0%, 100%': { boxShadow: '0 0 0 0 rgba(255, 215, 0, 0.4)' },
          '50%': { boxShadow: '0 0 0 10px rgba(255, 215, 0, 0)' },
        },
        bounceSlight: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%': { transform: 'translateY(-5px)' },
        },
        shimmer: {
          '0%': { transform: 'translateX(-100%)' },
          '100%': { transform: 'translateX(100%)' },
        },
        pulse: {
          '0%, 100%': { transform: 'scale(1)', opacity: '1' },
          '50%': { transform: 'scale(1.05)', opacity: '0.8' },
        },
        goldenPulse: {
          '0%': { boxShadow: '0 0 0 0 rgba(255, 215, 0, 0.4)' },
          '70%': { boxShadow: '0 0 0 10px rgba(255, 215, 0, 0)' },
          '100%': { boxShadow: '0 0 0 0 rgba(255, 215, 0, 0)' },
        },
        float: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%': { transform: 'translateY(-5px)' },
        },
      },
      transitionProperty: {
        'height': 'height',
        'spacing': 'margin, padding',
      },
      zIndex: {
        '60': '60',
        '70': '70',
        '80': '80',
        '90': '90',
        '100': '100',
      },
      backgroundImage: {
        'hero-gradient': 'linear-gradient(to right, var(--tw-gradient-stops))',
        'golden-shimmer': 'linear-gradient(90deg, transparent, rgba(255, 215, 0, 0.3), transparent)',
      },
    },
  },
  plugins: [
    // Custom utilities plugin
    plugin(({ addUtilities }) => {
      const newUtilities = {
        '.text-shadow': {
          textShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
        },
        '.text-shadow-md': {
          textShadow: '0 4px 8px rgba(0, 0, 0, 0.12), 0 2px 4px rgba(0, 0, 0, 0.08)',
        },
        '.text-shadow-none': {
          textShadow: 'none',
        },
        '.bg-blur': {
          backdropFilter: 'blur(8px)',
        },
        '.bg-blur-sm': {
          backdropFilter: 'blur(4px)',
        },
        '.bg-blur-lg': {
          backdropFilter: 'blur(16px)',
        },
        '.golden-glow': {
          boxShadow: '0 0 15px rgba(255, 215, 0, 0.5)',
        },
        '.golden-glow-sm': {
          boxShadow: '0 0 10px rgba(255, 215, 0, 0.3)',
        },
        '.golden-glow-lg': {
          boxShadow: '0 0 20px rgba(255, 215, 0, 0.7)',
        },
        '.shimmer-overlay': {
          position: 'relative',
          overflow: 'hidden',
          '&::after': {
            content: '""',
            position: 'absolute',
            top: '0',
            right: '0',
            bottom: '0',
            left: '0',
            transform: 'translateX(-100%)',
            backgroundImage: 'linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent)',
            animation: 'shimmer 2.5s infinite',
          },
        },
        '.golden-shimmer-overlay': {
          position: 'relative',
          overflow: 'hidden',
          '&::after': {
            content: '""',
            position: 'absolute',
            top: '0',
            right: '0',
            bottom: '0',
            left: '0',
            transform: 'translateX(-100%)',
            backgroundImage: 'linear-gradient(90deg, transparent, rgba(255, 215, 0, 0.3), transparent)',
            animation: 'shimmer 2.5s infinite',
          },
        },
      }
      addUtilities(newUtilities)
    }),
    // Animation delay plugin
    plugin(({ matchUtilities, theme }) => {
      matchUtilities(
        {
          'animate-delay': (value) => ({
            animationDelay: value,
          }),
        },
        { values: theme('transitionDelay') }
      )
    }),
    // Scrollbar customization plugin
    plugin(({ addUtilities }) => {
      const scrollbarUtilities = {
        '.scrollbar-thin': {
          scrollbarWidth: 'thin',
          '&::-webkit-scrollbar': {
            width: '6px',
            height: '6px',
          },
        },
        '.scrollbar-custom': {
          scrollbarWidth: 'thin',
          '&::-webkit-scrollbar': {
            width: '6px',
            height: '6px',
          },
          '&::-webkit-scrollbar-track': {
            background: '#f1f1f1',
            borderRadius: '100vh',
          },
          '&::-webkit-scrollbar-thumb': {
            background: '#c1c1c1',
            borderRadius: '100vh',
          },
          '&::-webkit-scrollbar-thumb:hover': {
            background: '#a8a8a8',
          },
        },
        '.scrollbar-golden': {
          scrollbarWidth: 'thin',
          '&::-webkit-scrollbar': {
            width: '6px',
            height: '6px',
          },
          '&::-webkit-scrollbar-track': {
            background: '#f1f1f1',
            borderRadius: '100vh',
          },
          '&::-webkit-scrollbar-thumb': {
            background: 'linear-gradient(#FFD700, #E6C200)',
            borderRadius: '100vh',
          },
          '&::-webkit-scrollbar-thumb:hover': {
            background: '#E6C200',
          },
        },
      }
      addUtilities(scrollbarUtilities)
    }),
  ],
} satisfies Config