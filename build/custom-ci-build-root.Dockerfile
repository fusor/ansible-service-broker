FROM openshift/origin-release:golang-1.13

RUN yum install -y epel-release \
    && yum install -y python-devel python-pip gcc

RUN pip install -U setuptools wheel && pip install -U more-itertools==7.0.0 molecule==2.22.0 jmespath openshift

RUN chmod g+rw /etc/passwd
