import React from 'react'
import ButtonWithIcon from '../../button_with_icon'

interface Props {
  title: string
  pub_date: string
  duration: number
}

const Episode = ({ title, pubDate, duration }) => {
  return (
    <>
      <div className="flex justify-between my-2">
        <div className="flex-auto">
          <h4 className="text-sm font-medium">{title}</h4>
          <span className="text-xs">
            1 hr 2 min
            <span className="mx-2 font-extrabold">&middot;</span>2 days ago
          </span>
        </div>
        <ButtonWithIcon className="w-5" icon="play" />
      </div>
      <hr className="my-3" />
    </>
  )
}

export default Episode
