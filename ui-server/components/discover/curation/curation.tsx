import Grid from 'components/grid'
import Link from 'next/link'
import React from 'react'
import { Curation, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  curation: Curation
  podcasts: Podcast[]
}

export interface OwnProps {
  curationId: string
}

interface Props extends StateToProps, OwnProps {}

const CurationView: React.SFC<Props> = (props) => {
  const { curation, podcasts } = props

  return (
    <div>
      <div className="pb-5">
        <h3 className="text-xl text-gray-800">{curation.title}</h3>
      </div>
      <Grid rows={1} cols={7} className="mb-4 pb-8" classNameChild="w-36 mx-2">
        {podcasts.map((podcast) => (
          <Link
            href={{ pathname: '/podcasts', query: { podcastId: podcast.id } }}
            as={`/podcasts/${podcast.id}`}
            key={podcast.id}
          >
            <div className="w-full cursor-pointer">
              <img
                className="w-full h-auto flex-none object-contain rounded-lg border"
                src={getImageUrl(podcast.id, 'md')}
              />
              <p className="text-xs tracking-wide leading-tight text-gray-800 mt-2 mb-1 line-clamp-2">
                {podcast.title}
              </p>
              <p className="text-xs tracking-tigher text-gray-600 truncate">
                {`by ${podcast.author}`}
              </p>
            </div>
          </Link>
        ))}
      </Grid>
    </div>
  )
}

export default CurationView
