import React from 'react'

interface Props {
  title: string
  author: string
  description: string
}

const Podcast: React.SFC<Props> = ({ title, author, description }) => (
  <div className="flex mb-8">
    <img
      className="lg:h-48 lg:w-48 h-32 w-32 flex-none object-cover object-center rounded"
      src="https://is1-ssl.mzstatic.com/image/thumb/Podcasts113/v4/37/f9/c4/37f9c4c9-c628-bb4f-37b1-1b6fef8f18a7/mza_7296818625298515281.jpeg/400x400.jpg"
    />
    <div className="flex flex-col lg:px-5 px-3">
      <h2 className="text-2xl text-gray-900">{title}</h2>
      <h3 className="text-lg text-gray-800 leading-relaxed">{author}</h3>
      <span className="lg:block hidden mt-1 text-sm text-gray-800 ">
        {description}
      </span>
    </div>
  </div>
)

export default Podcast
