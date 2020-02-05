import PageLayout from 'components/page_layout'
import { AboutPageSeo } from 'components/seo'
import React from 'react'

export default class AboutPage extends React.Component<{}> {
  render() {
    return (
      <>
        <AboutPageSeo />
        <PageLayout>
          <div>
            <h1 className="pt-3 pb-1 text-xl tracking-wide text-gray-800 font-medium">
              {'About'}
            </h1>
            <hr className="mb-3" />
            <div className="text-gray-700 tracking-wide font-normal">
              <p className="text-gray-700 text-base tracking-wide">
                {'Hi, My name is '}
                <a
                  href="https://www.linkedin.com/in/varmamsp/"
                  target="_blank"
                  className="underline text-purple-700"
                >
                  {'Pavan Varma'}
                </a>
                {". I'am the developer of Phenopod and nice to meet you."}
              </p>
              <br />
              <p className="text-gray-700 text-base tracking-wide">
                {
                  'Phenopod is a online podcast player and podcast directory, with over 638,417 podcasts and 21,708,559 episodes.'
                }
              </p>
              <p className="mt-1 text-gray-700 text-base tracking-wide">
                {
                  'The podcast ecosystem has matured over last few years and there is a podcast for just about every topic imaginable.'
                }
              </p>
              <p className="mt-1 text-gray-700 text-base tracking-wide">
                {
                  'Using Phenopod you can search episodes, subscribe to podcasts and create playlists of your favourite episodes.'
                }
              </p>
              <br />
              <p className="text-gray-700">
                {
                  'Phenopod is in very early development stage right now and I would love to hear your feedback.'
                }
                <br />
                {'Please reach out to me by one of the following means.'}
              </p>
              <ul className="list-disc pl-8 pt-2">
                <li className="my-1">
                  {'Twitter '}
                  <a
                    href="https://twitter.com/phenopod"
                    target="_blank"
                    className="underline text-purple-700"
                  >
                    {'@phenopod'}
                  </a>
                  {'   and   '}
                  <a
                    href="https://twitter.com/VarmaMSP"
                    target="_blank"
                    className="underline text-purple-700"
                  >
                    {'@varmamsp'}
                  </a>
                </li>
                <li className="my-1">
                  {'Reddit '}
                  <a
                    href="https://www.reddit.com/r/phenopod/"
                    target="_blank"
                    className="underline text-purple-700"
                  >
                    {'r/phenopod'}
                  </a>
                </li>
                <li className="my-1">
                  {'Email '}
                  <a
                    href="mailto:hello@phenopod.com"
                    className="underline text-purple-700"
                  >
                    {'hello@phenopod.com'}
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </PageLayout>
      </>
    )
  }
}
