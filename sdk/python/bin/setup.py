# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "3.3.2"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "stackmgmt Pulumi Package - Development Version"


setup(name='pequod_stackmgmt',
      python_requires='>=3.7',
      version=VERSION,
      long_description=readme(),
      long_description_content_type='text/markdown',
      packages=find_packages(),
      package_data={
          'pequod_stackmgmt': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=3.0.0,<4.0.0',
          'pulumi-command>=0.0.0,<1.0.0',
          'pulumi-pulumiservice>=0.0.0,<1.0.0',
          'semver>=2.8.1'
      ],
      zip_safe=False)
