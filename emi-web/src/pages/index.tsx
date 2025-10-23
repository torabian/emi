import type { ReactNode } from "react";
import clsx from "clsx";
import Link from "@docusaurus/Link";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Layout from "@theme/Layout";
import Heading from "@theme/Heading";

import styles from "./index.module.css";

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx("hero hero--primary", styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/intro"
            style={{ margin: "0 10px" }}
          >
            Learn more
          </Link>
          <a
            className="button button--secondary button--lg"
            href="/emi/playground"
          >
            Playground
          </a>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Emi Compiler`}
      description="Emi compiler documents and playground"
    >
      <HomepageHeader />
      <div className="container" style={{ textAlign: "center" }}>
        <div style={{ margin: "40px 0" }}>
          <strong>
            You can learn about the emi in this playlist on youtube:{" "}
            <a href="https://www.youtube.com/playlist?list=PLT2M-bOOx8oyAsIg4jdM-ErmlZ2MqE_SJ">
              Emi Compiler Training
            </a>
          </strong>
        </div>

        <br />

        <iframe
          style={{ width: "100%", height: "50vh", marginBottom: "50px" }}
          src="https://www.youtube.com/embed/p_hkbarIQjM"
          title="YouTube video player"
          frameBorder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
          referrerPolicy="strict-origin-when-cross-origin"
          allowFullScreen
        ></iframe>

        <iframe
          style={{ width: "100%", height: "50vh", marginBottom: "50px" }}
          src="https://www.youtube.com/embed/4JwsYdvOAy4"
          title="YouTube video player"
          frameBorder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
          referrerPolicy="strict-origin-when-cross-origin"
          allowFullScreen
        ></iframe>
      </div>
    </Layout>
  );
}
