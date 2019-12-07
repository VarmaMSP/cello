import React from 'react'
import { PodcastList } from 'types/app'

export interface StateToProps {
  categories: PodcastList[]
}

const Categories: React.FC<StateToProps> = ({ categories }) => {
  return (
    <div>
      <h2 className="text-xl text-gray-700">{'Browse by categories'}</h2>
      <hr className="my-2" />
      <ul>
        {categories.map((c) => (
          <li
            key={c.id}
            className="block flex items-center h-10 w-2/3 pl-6 rounded-full hover:bg-gray-200"
          >
            <span className="text-lg text-gray-900">{c.title}</span>
            &nbsp;&nbsp;
            <span className="text-sm text-gray-700 tracking-wider">
              {c.subTitle}
            </span>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default Categories
