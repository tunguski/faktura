FROM tunguski/go-dev
MAINTAINER Marek Romanowski

#ENV DEBIAN_FRONTEND noninteractive
USER root

RUN apt-get update -q && \
    apt-get install -qy texlive-full python-pygments gnuplot make git && \
    rm -rf /var/lib/apt/lists/*

#RUN apt-get update
#RUN apt-get install -y build-essential wget libfontconfig1 apt-transport-https
#RUN apt-get install -y tex-common tex-gyre texlive-base texlive-binaries texlive-fonts-recommended texlive-lang-polish texlive-latex-base texlive-latex-recommended texlive
#RUN rm -vrf /var/lib/apt/lists/*

# Install TexLive with scheme-basic
#RUN wget http://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz; \
#	mkdir /install-tl-unx; \
#	tar -xvf install-tl-unx.tar.gz -C /install-tl-unx --strip-components=1; \
#    echo "selected_scheme scheme-basic" >> /install-tl-unx/texlive.profile; \
#	/install-tl-unx/install-tl -profile /install-tl-unx/texlive.profile; \
#    rm -r /install-tl-unx; \
#	rm install-tl-unx.tar.gz

#ENV PATH="/usr/local/texlive/2017/bin/x86_64-linux:${PATH}"

#ENV HOME /data
#WORKDIR /data

USER dev
ENV HOME /home/dev

# Install latex packages
#RUN tlmgr install latexmk

VOLUME ["/data"]
