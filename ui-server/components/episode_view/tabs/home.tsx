import { Episode } from 'models'
import React, { useEffect, useRef } from 'react'

interface OwnProps {
  episode: Episode
}

const HomeTab: React.FC<OwnProps> = ({ episode }) => {
  const ref = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (ref.current) {
      const a = ref.current.getElementsByTagName('a')
      for (let i = 0; i < a.length; ++i) {
        a[i].setAttribute('target', '_blank')
      }

      const img = ref.current.getElementsByTagName('img')
      for (let i = 0; i < img.length; ++i) {
        img[i].remove()
      }
    }
  })

  return (
    <div>
      <div
        ref={ref}
        className="external-html lg:pr-10 text-sm text-gray-800 font-medium tracking-wide leading-relaxed"
        dangerouslySetInnerHTML={{ __html: episode.description }}
      />
    </div>
  )
}

export default HomeTab
