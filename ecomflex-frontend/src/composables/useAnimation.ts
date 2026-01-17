import { ref, onMounted, onUnmounted } from 'vue'

export type AnimationTiming = 'linear' | 'ease' | 'ease-in' | 'ease-out' | 'ease-in-out'
export type AnimationDirection = 'normal' | 'reverse' | 'alternate' | 'alternate-reverse'
export type AnimationFillMode = 'none' | 'forwards' | 'backwards' | 'both'

export interface AnimationOptions {
  timing?: AnimationTiming
  duration?: number
  delay?: number
  iterations?: number | 'infinite'
  direction?: AnimationDirection
  fillMode?: AnimationFillMode
}

export interface AnimationInstance {
  element: HTMLElement
  animation: Animation
  play: () => void
  pause: () => void
  cancel: () => void
  finish: () => void
  reverse: () => void
  isPending: () => boolean
  isRunning: () => boolean
  isPaused: () => boolean
  isFinished: () => boolean
}

export interface CountUpOptions {
  duration?: number
  decimals?: number
  startValue?: number
  delay?: number
  easing?: AnimationTiming
}

export function useAnimation() {
  const animations = ref<AnimationInstance[]>([])

  // Cleanup on component unmount
  onUnmounted(() => {
    animations.value.forEach(instance => {
      instance.cancel()
    })
    animations.value = []
  })

  /**
   * Fade in animation
   */
  const fadeIn = (
    element: HTMLElement,
    options: AnimationOptions = {}
  ): AnimationInstance => {
    const defaultOptions: AnimationOptions = {
      timing: 'ease-out',
      duration: 300,
      delay: 0,
      iterations: 1,
      direction: 'normal',
      fillMode: 'forwards'
    }
    
    const mergedOptions = { ...defaultOptions, ...options }
    
    const keyframes = [
      { opacity: 0 },
      { opacity: 1 }
    ]
    
    const animation = element.animate(keyframes, {
      duration: mergedOptions.duration,
      delay: mergedOptions.delay,
      easing: mergedOptions.timing,
      iterations: mergedOptions.iterations,
      direction: mergedOptions.direction,
      fill: mergedOptions.fillMode
    })
    
    const instance: AnimationInstance = {
      element,
      animation,
      play: () => animation.play(),
      pause: () => animation.pause(),
      cancel: () => animation.cancel(),
      finish: () => animation.finish(),
      reverse: () => animation.reverse(),
      isPending: () => animation.pending,
      isRunning: () => animation.playState === 'running',
      isPaused: () => animation.playState === 'paused',
      isFinished: () => animation.playState === 'finished'
    }
    
    animations.value.push(instance)
    return instance
  }
  
  /**
   * Fade out animation
   */
  const fadeOut = (
    element: HTMLElement,
    options: AnimationOptions = {}
  ): AnimationInstance => {
    const defaultOptions: AnimationOptions = {
      timing: 'ease-in',
      duration: 300,
      delay: 0,
      iterations: 1,
      direction: 'normal',
      fillMode: 'forwards'
    }
    
    const mergedOptions = { ...defaultOptions, ...options }
    
    const keyframes = [
      { opacity: 1 },
      { opacity: 0 }
    ]
    
    const animation = element.animate(keyframes, {
      duration: mergedOptions.duration,
      delay: mergedOptions.delay,
      easing: mergedOptions.timing,
      iterations: mergedOptions.iterations,
      direction: mergedOptions.direction,
      fill: mergedOptions.fillMode
    })
    
    const instance: AnimationInstance = {
      element,
      animation,
      play: () => animation.play(),
      pause: () => animation.pause(),
      cancel: () => animation.cancel(),
      finish: () => animation.finish(),
      reverse: () => animation.reverse(),
      isPending: () => animation.pending,
      isRunning: () => animation.playState === 'running',
      isPaused: () => animation.playState === 'paused',
      isFinished: () => animation.playState === 'finished'
    }
    
    animations.value.push(instance)
    return instance
  }
  
  /**
   * Slide in animation (from direction)
   */
  const slideIn = (
    element: HTMLElement,
    direction: 'top' | 'right' | 'bottom' | 'left' = 'bottom',
    distance: string = '20px',
    options: AnimationOptions = {}
  ): AnimationInstance => {
    const defaultOptions: AnimationOptions = {
      timing: 'ease-out',
      duration: 300,
      delay: 0,
      iterations: 1,
      direction: 'normal',
      fillMode: 'forwards'
    }
    
    const mergedOptions = { ...defaultOptions, ...options }
    
    let fromTransform = ''
    
    switch (direction) {
      case 'top':
        fromTransform = `translateY(-${distance})`
        break
      case 'right':
        fromTransform = `translateX(${distance})`
        break
      case 'bottom':
        fromTransform = `translateY(${distance})`
        break
      case 'left':
        fromTransform = `translateX(-${distance})`
        break
    }
    
    const keyframes = [
      { opacity: 0, transform: fromTransform },
      { opacity: 1, transform: 'translate(0)' }
    ]
    
    const animation = element.animate(keyframes, {
      duration: mergedOptions.duration,
      delay: mergedOptions.delay,
      easing: mergedOptions.timing,
      iterations: mergedOptions.iterations,
      direction: mergedOptions.direction,
      fill: mergedOptions.fillMode
    })
    
    const instance: AnimationInstance = {
      element,
      animation,
      play: () => animation.play(),
      pause: () => animation.pause(),
      cancel: () => animation.cancel(),
      finish: () => animation.finish(),
      reverse: () => animation.reverse(),
      isPending: () => animation.pending,
      isRunning: () => animation.playState === 'running',
      isPaused: () => animation.playState === 'paused',
      isFinished: () => animation.playState === 'finished'
    }
    
    animations.value.push(instance)
    return instance
  }
  
  /**
   * Scale animation
   */
  const scale = (
    element: HTMLElement,
    fromScale: number = 0.8,
    toScale: number = 1,
    options: AnimationOptions = {}
  ): AnimationInstance => {
    const defaultOptions: AnimationOptions = {
      timing: 'ease-out',
      duration: 300,
      delay: 0,
      iterations: 1,
      direction: 'normal',
      fillMode: 'forwards'
    }
    
    const mergedOptions = { ...defaultOptions, ...options }
    
    const keyframes = [
      { opacity: fromScale < 1 ? 0 : 1, transform: `scale(${fromScale})` },
      { opacity: 1, transform: `scale(${toScale})` }
    ]
    
    const animation = element.animate(keyframes, {
      duration: mergedOptions.duration,
      delay: mergedOptions.delay,
      easing: mergedOptions.timing,
      iterations: mergedOptions.iterations,
      direction: mergedOptions.direction,
      fill: mergedOptions.fillMode
    })
    
    const instance: AnimationInstance = {
      element,
      animation,
      play: () => animation.play(),
      pause: () => animation.pause(),
      cancel: () => animation.cancel(),
      finish: () => animation.finish(),
      reverse: () => animation.reverse(),
      isPending: () => animation.pending,
      isRunning: () => animation.playState === 'running',
      isPaused: () => animation.playState === 'paused',
      isFinished: () => animation.playState === 'finished'
    }
    
    animations.value.push(instance)
    return instance
  }
  
  /**
   * Pulse animation (golden glow)
   */
  const pulseGlow = (
    element: HTMLElement,
    color: string = 'rgba(255, 215, 0, 0.5)', // Golden glow
    options: AnimationOptions = {}
  ): AnimationInstance => {
    const defaultOptions: AnimationOptions = {
      timing: 'ease-in-out',
      duration: 1500,
      delay: 0,
      iterations: 'infinite',
      direction: 'alternate',
      fillMode: 'none'
    }
    
    const mergedOptions = { ...defaultOptions, ...options }
    
    const keyframes = [
      { boxShadow: `0 0 0 0 ${color}` },
      { boxShadow: `0 0 0 10px transparent` }
    ]
    
    const animation = element.animate(keyframes, {
      duration: mergedOptions.duration,
      delay: mergedOptions.delay,
      easing: mergedOptions.timing,
      iterations: mergedOptions.iterations,
      direction: mergedOptions.direction,
      fill: mergedOptions.fillMode
    })
    
    const instance: AnimationInstance = {
      element,
      animation,
      play: () => animation.play(),
      pause: () => animation.pause(),
      cancel: () => animation.cancel(),
      finish: () => animation.finish(),
      reverse: () => animation.reverse(),
      isPending: () => animation.pending,
      isRunning: () => animation.playState === 'running',
      isPaused: () => animation.playState === 'paused',
      isFinished: () => animation.playState === 'finished'
    }
    
    animations.value.push(instance)
    return instance
  }
  
  /**
   * Count up animation (for numbers)
   */
  const countUp = (
    element: HTMLElement,
    targetValue: number,
    options: CountUpOptions = {}
  ) => {
    const defaultOptions: CountUpOptions = {
      duration: 1500,
      decimals: 0,
      startValue: 0,
      delay: 0,
      easing: 'ease-out'
    }
    
    const mergedOptions = { ...defaultOptions, ...options }
    
    // Store the original text for later
    const originalText = element.textContent || ''
    
    // Create an animation
    const startValue = mergedOptions.startValue || 0
    const totalChange = targetValue - startValue
    
    const updateValue = (progress: number) => {
      const currentValue = startValue + totalChange * progress
      element.textContent = currentValue.toFixed(mergedOptions.decimals)
    }
    
    // Timing function
    const getEasingProgress = (t: number): number => {
      switch (mergedOptions.easing) {
        case 'linear':
          return t
        case 'ease-in':
          return t * t
        case 'ease-out':
          return t * (2 - t)
        case 'ease-in-out':
          return t < 0.5 ? 2 * t * t : -1 + (4 - 2 * t) * t
        default: // ease
          return t < 0.5 ? 2 * t * t : -1 + (4 - 2 * t) * t
      }
    }
    
    // Animation start time
    let startTime: number | null = null
    let animationFrame: number | null = null
    
    const animate = (time: number) => {
      if (!startTime) {
        startTime = time
      }
      
      const elapsedTime = time - startTime
      const progress = Math.min(elapsedTime / mergedOptions.duration!, 1)
      const easedProgress = getEasingProgress(progress)
      
      updateValue(easedProgress)
      
      if (progress < 1) {
        animationFrame = requestAnimationFrame(animate)
      }
    }
    
    // Start the animation after delay
    setTimeout(() => {
      animationFrame = requestAnimationFrame(animate)
    }, mergedOptions.delay)
    
    // Return an object with control methods
    return {
      stop: () => {
        if (animationFrame) {
          cancelAnimationFrame(animationFrame)
          animationFrame = null
        }
      },
      reset: () => {
        if (animationFrame) {
          cancelAnimationFrame(animationFrame)
          animationFrame = null
        }
        element.textContent = originalText
      }
    }
  }
  
  /**
   * Button click ripple effect
   */
  const buttonRipple = (
    event: MouseEvent,
    options: {
      color?: string,
      duration?: number
    } = {}
  ) => {
    const button = event.currentTarget as HTMLElement
    if (!button) return
    
    const defaultOptions = {
      color: 'rgba(255, 215, 0, 0.3)', // Golden ripple
      duration: 600
    }
    
    const mergedOptions = { ...defaultOptions, ...options }
    
    // Get button dimensions and position
    const rect = button.getBoundingClientRect()
    
    // Calculate ripple size
    const size = Math.max(rect.width, rect.height) * 2
    
    // Calculate ripple position
    const x = event.clientX - rect.left
    const y = event.clientY - rect.top
    
    // Create ripple element
    const ripple = document.createElement('span')
    ripple.style.position = 'absolute'
    ripple.style.width = '0'
    ripple.style.height = '0'
    ripple.style.borderRadius = '50%'
    ripple.style.transform = 'translate(-50%, -50%)'
    ripple.style.backgroundColor = mergedOptions.color
    ripple.style.left = `${x}px`
    ripple.style.top = `${y}px`
    
    // Ensure button has relative position for absolute positioning of ripple
    const computedStyle = window.getComputedStyle(button)
    if (computedStyle.position === 'static') {
      button.style.position = 'relative'
    }
    
    button.style.overflow = 'hidden'
    button.appendChild(ripple)
    
    // Animate the ripple
    const animation = ripple.animate(
      [
        { width: '0', height: '0', opacity: 0.5 },
        { width: `${size}px`, height: `${size}px`, opacity: 0 }
      ],
      {
        duration: mergedOptions.duration,
        easing: 'ease-out',
        fill: 'forwards'
      }
    )
    
    // Remove ripple after animation completes
    animation.onfinish = () => {
      ripple.remove()
    }
  }

  return {
    fadeIn,
    fadeOut,
    slideIn,
    scale,
    pulseGlow,
    countUp,
    buttonRipple
  }
}