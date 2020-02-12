import React, { useEffect, useRef } from 'react'
import { Episode } from 'models'

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
      <h2 className="font-medium tracking-wider mb-2">{'Description'}</h2>
      <div
        ref={ref}
        className="external-html lg:pr-16 text-sm leading-relaxed tracking-wide text-gray-900"
        dangerouslySetInnerHTML={{ __html: episode.description }}
      />
    </div>
  )
}

export default HomeTab
