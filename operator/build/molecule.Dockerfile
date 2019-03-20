FROM centos:7

RUN mkdir -p /home/molecule

ENV HOME=/home/molecule
ENV PYTHONUSERBASE=${HOME}/.bin
RUN mkdir -p ${PYTHONUSERBASE}
ENV PATH="${HOME}/.bin:${PATH}"

RUN yum install -y epel-release \
    && yum install -y python-devel python-pip gcc

RUN pip install --user -U setuptools && pip install --user molecule==2.20.0.0a2 jmespath openshift

COPY . ${HOME}
